package model

import "github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"

type (
	VerifyRequest struct {
		WsContext global.WsContext `json:"WSContext"`
		Token     string           `json:"Token"`
		RefNum    string           `json:"RefNum"`
	}

	VerifyResponse struct {
		Result string `json:"Result"`
		RefNum string `json:"RefNum"`
		Amount uint64 `json:"Amount"`
	}
)

func GetVerifyTransactionRequestBody(token, refNum string) VerifyRequest {
	return VerifyRequest{
		WsContext: global.GetWsContext(),
		Token:     token,
		RefNum:    refNum,
	}
}

func (res VerifyResponse) GetResult() string {
	return res.Result
}
