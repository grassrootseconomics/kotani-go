package kotani

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type (
	CreateMobileMoneyCustomerBody struct {
		CountryCode CountryCode `json:"country_code"`
		PhoneNumber string      `json:"phone_number"`
		Network     Network     `json:"network"`
	}

	UpdateMobileMoneyCustomerBody struct {
		CountryCode CountryCode `json:"country_code"`
		Network     Network     `json:"network"`
		AccountName string      `json:"account_name"`
		IDNumber    string      `json:"id_number"`
		IDType      IDType      `json:"id_type"`
	}

	MobileMoneyCustomerResp struct {
		CommonResponseFields
		Data struct {
			ID          string      `json:"id"`
			PhoneNumber string      `json:"phone_number"`
			CountryCode CountryCode `json:"country_code"`
			Network     Network     `json:"network"`
			CustomerKey string      `json:"customer_key"`
			AccountName string      `json:"account_name"`
			Integrator  string      `json:"integrator"`
		} `json:"data"`
	}
)

const customerPath = "/customer/"

func (kc *KotaniClient) CreateMobileMoneyCustomer(ctx context.Context, input CreateMobileMoneyCustomerBody) (MobileMoneyCustomerResp, error) {
	mobileMoneyCustomerResp := MobileMoneyCustomerResp{}

	jsonRequestBody, err := json.Marshal(&input)
	if err != nil {
		return mobileMoneyCustomerResp, err
	}

	resp, err := kc.requestWithCtx(ctx, http.MethodPost, kc.endpoint+customerPath+"mobile-money", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return mobileMoneyCustomerResp, err
	}

	if err := parseResponse(resp, &mobileMoneyCustomerResp); err != nil {
		return mobileMoneyCustomerResp, err
	}

	return mobileMoneyCustomerResp, nil
}

func (kc *KotaniClient) UpdateMobileMoneyCustomer(ctx context.Context, customerKey string, input UpdateMobileMoneyCustomerBody) (MobileMoneyCustomerResp, error) {
	mobileMoneyCustomerResp := MobileMoneyCustomerResp{}

	jsonRequestBody, err := json.Marshal(&input)
	if err != nil {
		return mobileMoneyCustomerResp, err
	}

	resp, err := kc.requestWithCtx(ctx, http.MethodPatch, kc.endpoint+customerPath+"mobile-money/"+customerKey, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return mobileMoneyCustomerResp, err
	}

	if err := parseResponse(resp, &mobileMoneyCustomerResp); err != nil {
		return mobileMoneyCustomerResp, err
	}

	return mobileMoneyCustomerResp, nil
}

func (kc *KotaniClient) GetMobileMoneyCustomerByPhone(ctx context.Context, phoneNumber string) (MobileMoneyCustomerResp, error) {
	mobileMoneyCustomerResp := MobileMoneyCustomerResp{}

	resp, err := kc.requestWithCtx(ctx, http.MethodGet, kc.endpoint+customerPath+"mobile-money/phone/"+phoneNumber, nil)
	if err != nil {
		return mobileMoneyCustomerResp, err
	}

	if err := parseResponse(resp, &mobileMoneyCustomerResp); err != nil {
		return mobileMoneyCustomerResp, err
	}

	return mobileMoneyCustomerResp, nil
}
