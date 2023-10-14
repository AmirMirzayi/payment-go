package novinpay

import (
	"context"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/helper"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/model"
)

func login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.WaitingForResponseTime)
	defer cancel()

	txResult, err := helper.NewPostRequestWithContext[model.LoginResponse](
		ctx,
		global.LoginURI,
		model.GetLoginRequestBody(),
	)
	if err != nil {
		return "", err
	}

	return txResult.SessionId, nil
}
