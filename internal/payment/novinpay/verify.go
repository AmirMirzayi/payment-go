package novinpay

import (
	"context"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/helper"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/model"
)

func VerifyTransaction(token, refNum string) (string, uint64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.WaitingForResponseTime)
	defer cancel()

	txResult, err := helper.NewPostRequestWithContext[model.VerifyResponse](
		ctx,
		global.VerifyTransaction,
		model.GetVerifyTransactionRequestBody(token, refNum),
	)
	if err != nil {
		return "", 0, err
	}

	return txResult.RefNum, txResult.Amount, nil
}
