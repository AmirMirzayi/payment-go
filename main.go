package main

import "github.com/AmirMirzayi/payment-go/internal/payment/novinpay"

func main() {
	err := novinpay.
		NewBuyTransaction().
		SetAmount(12000).
		SetMobile("09105599950").
		Pay()

	if err != nil {
		panic(err)
	}
}
