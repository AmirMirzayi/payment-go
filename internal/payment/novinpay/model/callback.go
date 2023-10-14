package model

type PaymentCallback struct {
	State             string `json:"State"`
	ResNum            string `json:"ResNum"`
	RefNum            string `json:"RefNum"`
	CustomerRefNum    string `json:"CustomerRefNum"`
	MID               string `json:"MID"`
	Language          string `json:"Language"`
	CardHashPan       string `json:"CardHashPan"`
	CardMaskPan       string `json:"CardMaskPan"`
	GoodReferenceId   string `json:"GoodReferenceId"`
	MerchantData      string `json:"MerchantData"`
	TraceNo           string `json:"TraceNo"`
	Token             string `json:"token"`
	TransactionAmount string `json:"transactionAmount"`
	Email             string `json:"emailAddress"`
	Mobile            string `json:"mobileNo"`
}
