package helper

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/global"
	"github.com/AmirMirzayi/payment-go/internal/payment/novinpay/model"
)

func IsSuccessful(msg string) bool {
	return strings.ToLower(msg) == strings.ToLower(global.ResponseMsg[global.Success])
}

func GetResponseText(msg string) string {
	return global.Code_enum[msg].String()
}

func GetResponseEnum(msg string) global.Code {
	return global.Code_enum[msg]
}

func NewPostRequestWithContext[T model.Response](
	ctx context.Context,
	uri global.PaymentURI,
	requestBody any,
) (response T, err error) {
	body, err := json.Marshal(requestBody)
	if err != nil {
		return response, err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		string(uri),
		bytes.NewBuffer(body),
	)
	request.Header.Set("Content-Type", "application/json")

	httpResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		if err == context.DeadlineExceeded {
			return response, fmt.Errorf("پاسخی از سمت پذیرنده درگاه پرداخت در زمان مقتضی دریافت نشد.")
		}
		return response, err
	}

	defer httpResponse.Body.Close()

	err = json.NewDecoder(httpResponse.Body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("خطا در دریافت اطلاعات از پذیرنده درگاه پرداخت.")
	}

	if httpResponse.StatusCode != http.StatusOK {
		return response, fmt.Errorf("خطا در برقراری ارتباط با پذیرنده درگاه پرداخت.")
	}

	if !IsSuccessful(response.GetResult()) {
		return response, fmt.Errorf("خطا در سرویس درگاه پرداخت: %s", GetResponseText(response.GetResult()))
	}

	return response, nil
}
