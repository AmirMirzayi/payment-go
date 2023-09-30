package main

import "github.com/AmirMirzayi/payment-go/internal/payment/novinpay"

func main() {
	err := novinpay.
		NewBuyTransaction("123456", 1111).
		SetMobile("09105599950").
		Pay()

	if err != nil {
		panic(err)
	}
}
