package kotani

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type (
	WalletBody struct {
		Name     string   `json:"name"`
		Currency Currency `json:"currency"`
	}

	CreateWalletResponse struct {
		CommonResponseFields
		Data struct {
			Name           string `json:"name"`
			Type           string `json:"type"`
			Currency       string `json:"currency"`
			Integrator     string `json:"integrator"`
			Status         string `json:"status"`
			ID             string `json:"id"`
			Balance        int    `json:"balance"`
			DepositBalance int    `json:"deposit_balance"`
		} `json:"data"`
	}

	FiatWalletsResponse struct {
		CommonResponseFields
		Data []struct {
			Name           string `json:"name"`
			Type           string `json:"type"`
			Currency       string `json:"currency"`
			Integrator     string `json:"integrator"`
			Status         string `json:"status"`
			ID             string `json:"id"`
			Balance        int    `json:"balance"`
			DepositBalance int    `json:"deposit_balance"`
		} `json:"data"`
	}
)

const walletPath = "/wallet/"

func (kc *KotaniClient) CreateFiatWallet(ctx context.Context, input WalletBody) (CreateWalletResponse, error) {
	createWalletResp := CreateWalletResponse{}

	jsonRequestBody, err := json.Marshal(&input)
	if err != nil {
		return createWalletResp, err
	}

	resp, err := kc.requestWithCtx(ctx, http.MethodPost, kc.endpoint+walletPath+"fiat", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return createWalletResp, err
	}

	if err := parseResponse(resp, &createWalletResp); err != nil {
		return createWalletResp, err
	}

	return createWalletResp, nil
}

func (kc *KotaniClient) GetIntegratorFiatWallets(ctx context.Context) (FiatWalletsResponse, error) {
	fiatWalletsResp := FiatWalletsResponse{}

	resp, err := kc.requestWithCtx(ctx, http.MethodGet, kc.endpoint+walletPath+"fiat", nil)
	if err != nil {
		return fiatWalletsResp, err
	}

	if err := parseResponse(resp, &fiatWalletsResp); err != nil {
		return fiatWalletsResp, err
	}

	return fiatWalletsResp, nil
}
