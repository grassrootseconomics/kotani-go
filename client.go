package kotani

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type KotaniClient struct {
	apiKey     string
	endpoint   string
	httpClient *http.Client
}

const (
	userAgent   = "kotani-go"
	contentType = "application/json"

	baseLiveEndpoint    = "https://api.kotanipay.io/api/v3"
	baseSandboxEndpoint = "https://sandbox-api.kotanipay.io/v3"
)

// New returns an instance of a Kotani client
func New(apiKey string, sandbox bool) *KotaniClient {
	kotaniClient := &KotaniClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	if sandbox {
		kotaniClient.endpoint = baseSandboxEndpoint
	} else {
		kotaniClient.endpoint = baseLiveEndpoint
	}

	return kotaniClient
}

// SetHTTPClient can be used to override the default client with a custom set one
func (kc *KotaniClient) SetHTTPClient(httpClient *http.Client) {
	kc.httpClient = httpClient
}

// setHeaders sets the headers required by the Fonbnk API
func (kc *KotaniClient) setHeaders(req *http.Request) (*http.Request, error) {
	req.Header.Set("Authorization", "Bearer "+kc.apiKey)
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", contentType)
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

// requestWithCtx builds the HTTP request
func (kc *KotaniClient) requestWithCtx(ctx context.Context, method string, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	return kc.do(req)
}

// do executes the built http request, setting appropriate headers
func (kc *KotaniClient) do(req *http.Request) (*http.Response, error) {
	builtRequest, err := kc.setHeaders(req)
	if err != nil {
		return nil, err
	}

	return kc.httpClient.Do(builtRequest)
}

// parseResponse is a general utility to decode JSON responses correctly
func parseResponse(resp *http.Response, target interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("kotani server error: code=%s: response_body=%s", resp.Status, string(b))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
