package global

import "time"

const WaitingForResponseTime time.Duration = 5 * time.Second

type Code uint8

const (
	Success Code = iota + 1
	InvalidUserOrPass
	InvalidSourceIp
	InvalidData
)

const (
	PaidStatus      string = "OK"
	CancelledStatus string = "Canceled By User"
)

type PaymentURI string

const (
	PaymentCallbackURL string = "https://dizintork.ir:4499/TestPayment/"

	PaymentURL string = "https://pna.shaparak.ir/_ipgw_/payment/"

	baseURI                       PaymentURI = "https://pna.shaparak.ir/ref-payment2/RestServices/mts"
	LoginURI                      PaymentURI = baseURI + "/merchantLogin/"
	GenerateTransactionDataToSign PaymentURI = baseURI + "/generateTransactionDataToSign/"
	GenerateSignedDataToken       PaymentURI = baseURI + "/generateSignedDataToken/"
	VerifyTransaction             PaymentURI = baseURI + "/verifyMerchantTrans/"
)

type WsContext struct {
	UserId   string `json:"UserId"`
	Password string `json:"Password"`
}
