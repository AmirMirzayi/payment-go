package global

type Code uint8

const (
	Success Code = iota + 1
	InvalidUserOrPass
	InvalidSourceIp
	InvalidData
)

type PaymentURI string

const (
	PaymentCallback string = "https://dizintork.ir:4499/TestPayment/"

	PaymentURL string = "https://pna.shaparak.ir/_ipgw_/payment/"

	baseURI                       PaymentURI = "https://pna.shaparak.ir/ref-payment2/RestServices/mts"
	LoginURI                      PaymentURI = baseURI + "/merchantLogin/"
	GenerateTransactionDataToSign PaymentURI = baseURI + "/generateTransactionDataToSign/"
	GenerateSignedDataToken       PaymentURI = baseURI + "/generateSignedDataToken/"
)

type WsContext struct {
	UserId   string `json:"UserId"`
	Password string `json:"Password"`
}
