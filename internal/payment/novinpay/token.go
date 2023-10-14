package novinpay

import (
	"context"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/helper"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/model"
)

func generateSignedDataTokenRequestBody(signature, uniqueId string) model.GenerateSignedDataTokenRequest {
	return model.GenerateSignedDataTokenRequest{
		WsContext: global.GetWsContext(),
		Signature: signature,
		UniqueId:  uniqueId,
	}
}

func generateSignedDataToken(signature, uniqueId string) (model.GenerateSignedDataTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), global.WaitingForResponseTime)
	defer cancel()

	txResult, err := helper.NewPostRequestWithContext[model.GenerateSignedDataTokenResponse](
		ctx,
		global.GenerateSignedDataToken,
		generateSignedDataTokenRequestBody(signature, uniqueId),
	)
	if err != nil {
		return txResult, err
	}

	return txResult, nil
}
