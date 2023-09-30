package model

type Verify struct {
	Result string `json:"Result"`
	RefNum string `json:"RefNum"`
	Amount uint64 `json:"Amount"`
}
