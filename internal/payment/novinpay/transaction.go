package novinpay

import (
	"context"
	"time"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/helper"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/model"
)

func generateTransactionDataToSign(tx *BuyTransaction) (model.GenerateTransactionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	txResult, err := helper.NewPostRequestWithContext[model.GenerateTransactionResponse](
		ctx,
		global.GenerateTransactionDataToSign,
		model.GetGenerateTransactionRequestBody(tx.GetAmount()),
	)
	if err != nil {
		return txResult, err
	}

	return txResult, nil
}
