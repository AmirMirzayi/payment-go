package model

import (
	"os"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
)

type (
	GenerateTransactionRequest struct {
		WsContext              global.WsContext `json:"WSContext"`
		TransType              string           `json:"TransType"`
		ReserveNum             string           `json:"ReserveNum"`
		MerchantId             string           `json:"MerchantId"`
		TerminalId             string           `json:"TerminalId"`
		Amount                 uint64           `json:"Amount"`
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

func GetGenerateTransactionRequestBody(amount uint64, txId, orderId, mobile, email string) GenerateTransactionRequest {
	return GenerateTransactionRequest{
		WsContext:              global.GetWsContext(),
		TransType:              "EN_GOODS",
		ReserveNum:             orderId,
		MerchantId:             os.Getenv("NOVIN_PAY_MERCHANT_ID"),
		TerminalId:             os.Getenv("NOVIN_PAY_TERMINAL_ID"),
		Amount:                 amount,
		ProductId:              orderId,
		GoodsReferenceID:       txId,
		MerchatGoodReferenceID: txId,
		MobileNo:               mobile,
		Email:                  email,
		RedirectUrl:            global.PaymentCallbackURL,
	}
}

func (res GenerateTransactionResponse) GetResult() string {
	return res.Result
}
