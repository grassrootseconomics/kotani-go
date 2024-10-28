package kotani

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"math"
	"net/http"
)

type (
	OfframpCustomer struct {
		PhoneNumber string  `json:"phoneNumber"`
		AccountName string  `json:"accountName"`
		Network     Network `json:"networkProvider"`
	}

	OfframpRequestBody struct {
		Chain       Chain           `json:"chain"`
		Token       Token           `json:"token"`
		Currency    Currency        `json:"currency"`
		Amount      float64         `json:"cryptoAmount"`
		FromAddress string          `json:"senderAddress"`
		ReferenceID string          `json:"referenceId"`
		Customer    OfframpCustomer `json:"mobileMoneyReceiver"`
	}

	OfframpResp struct {
		CommonResponseFields
		Data struct {
			ReferenceID           string  `json:"referenceId"`
			FiatAmount            float64 `json:"fiatAmount"`
			FiatTransactionAmount float64 `json:"fiatTransactionAmount"`
			CryptoAmount          float64 `json:"cryptoAmount"`
			FiatCurrency          string  `json:"fiatCurrency"`
			CustomerKey           string  `json:"customerKey"`
			FiatWalletID          string  `json:"fiatWalletId"`
			SenderAddress         string  `json:"senderAddress"`
			TransactionHash       string  `json:"transactionHash"`
			TransactionHashAmount float64 `json:"transactionHashAmount"`
			Status                string  `json:"status"`
			OnchainStatus         string  `json:"onchainStatus"`
			Rate                  struct {
			} `json:"rate"`
			EscrowAddress string `json:"escrowAddress"`
		} `json:"data"`
	}
)

const offrampPath = "/offramp/"

var ErrNegativeAmount = errors.New("amount must be greater than 0")

func (kc *KotaniClient) Offramp(ctx context.Context, input OfframpRequestBody) (OfframpResp, error) {
	offrampResp := OfframpResp{}

	if math.Signbit(input.Amount) {
		return offrampResp, ErrNegativeAmount
	}

	jsonRequestBody, err := json.Marshal(&input)
	if err != nil {
		return offrampResp, err
	}

	resp, err := kc.requestWithCtx(ctx, http.MethodPost, kc.endpoint+offrampPath, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return offrampResp, err
	}

	if err := parseResponse(resp, &offrampResp); err != nil {
		return offrampResp, err
	}

	return offrampResp, nil
}

func (kc *KotaniClient) GetOfframpStatus(ctx context.Context, referenceID string) (OfframpResp, error) {
	offrampStatusResp := OfframpResp{}

	resp, err := kc.requestWithCtx(ctx, http.MethodGet, kc.endpoint+offrampPath+referenceID, nil)
	if err != nil {
		return offrampStatusResp, err
	}

	if err := parseResponse(resp, &offrampStatusResp); err != nil {
		return offrampStatusResp, err
	}

	return offrampStatusResp, nil
}
