package model

import (
	"time"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
)

type (
	GenerateSignedDataTokenRequest struct {
		WsContext global.WsContext `json:"WSContext"`
		Signature string           `json:"Signature"`
		UniqueId  string           `json:"UniqueId"`
	}

	GenerateSignedDataTokenResponse struct {
		Result         string        `json:"Result"`
		ExpirationDate time.Duration `json:"ExpirationDate"`
		Token          string        `json:"Token"`
		ChannelId      string        `json:"ChannelId"`
		UserId         string        `json:"UserId"`
	}
)

func (res GenerateSignedDataTokenResponse) GetResult() string {
	return res.Result
}
