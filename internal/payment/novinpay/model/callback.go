package model

type PaymentCallback struct {
	State    string `json:"State"`
	ResNum   string `json:"ResNum"`
	RefNum   string `json:"RefNum"`
	CustomerRefNum string
	MID      string `json:"MID"`
	Language string `json:"Language"`
	CardHashPan string
	CardMaskPan string
	GoodReferenceId string
	MerchantData string
}
