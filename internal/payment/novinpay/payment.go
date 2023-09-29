package novinpay

import (
	"fmt"
	"log"
)

type BuyTransaction struct {
	amount    uint64
	email     string
	mobile    string
	reserveNo string
}

func NewBuyTransaction() *BuyTransaction {
	return &BuyTransaction{}
}

func (tx *BuyTransaction) GetAmount() uint64 {
	return tx.amount
}

func (tx *BuyTransaction) SetAmount(price uint64) *BuyTransaction {
	tx.amount = price
	return tx
}

func (tx *BuyTransaction) GetEmail() string {
	return tx.email
}

func (tx *BuyTransaction) SetEmail(email string) *BuyTransaction {
	tx.email = email
	return tx
}

func (tx *BuyTransaction) GetMobile() string {
	return tx.mobile
}

func (tx *BuyTransaction) SetMobile(mobile string) *BuyTransaction {
	tx.mobile = mobile
	return tx
}

func (tx *BuyTransaction) GetReserveNo() string {
	return tx.reserveNo
}

func (tx *BuyTransaction) SetReserveNo(reserveNo string) *BuyTransaction {
	tx.reserveNo = reserveNo
	return tx
}

func (tx *BuyTransaction) Pay() error {
	sessionId, err := login()
	if err != nil {
		return err
	}
	log.Println("1) login done.", sessionId)
	response, err := generateTransactionDataToSign(tx)
	if err != nil {
		return err
	}
	log.Println("2) transaction data generated.")

	signature, err := signToken(response.DataToSign)
	if err != nil {
		return err
	}
	log.Println("3) signature generated.")

	tokenResponse, err := generateSignedDataToken(signature, response.UniqueId)
	if err != nil {
		return err
	}
	log.Println("4) token generated.")
	fmt.Printf("#%v", tokenResponse)

	return nil
}
