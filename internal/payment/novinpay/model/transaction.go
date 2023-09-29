package model

import (
	"strconv"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
)

type (
	GenerateTransactionRequest struct {
		WsContext              global.WsContext `json:"WSContext"`
		TransType              string           `json:"TransType"`
		ReserveNum             string           `json:"ReserveNum"`
		MerchantId             string           `json:"MerchantId"`
		TerminalId             string           `json:"TerminalId"`
		Amount                 string           `json:"Amount"`
		ProductId              string           `json:"ProductId"`
		GoodsReferenceID       string           `json:"GoodsReferenceID"`
		MerchatGoodReferenceID string           `json:"MerchatGoodReferenceID"`
		MobileNo               string           `json:"MobileNo"`
		Email                  string           `json:"Email"`
		RedirectUrl            string           `json:"RedirectUrl"`
	}

	GenerateTransactionResponse struct {
		Result     string `json:"Result"`
		DataToSign string `json:"DataToSign"`
		UniqueId   string `json:"UniqueId"`
	}
)

func GetGenerateTransactionRequestBody(amount uint64) GenerateTransactionRequest {

	return GenerateTransactionRequest{
		WsContext:              global.GetWsContext(),
		TransType:              "EN_GOODS",
		ReserveNum:             "123456",
		MerchantId:             "01306251",
		TerminalId:             "",
		Amount:                 strconv.FormatUint(amount, 10),
		ProductId:              "",
		GoodsReferenceID:       "987654",
		MerchatGoodReferenceID: "111",
		MobileNo:               "09127125699",
		Email:                  "faraji80@yahoo.com",
		RedirectUrl:            global.PaymentCallback,
	}
}

func (res GenerateTransactionResponse) GetResult() string {
	return res.Result
}
