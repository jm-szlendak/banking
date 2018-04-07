package models

import (
	"crypto/md5"
	"fmt"
	"io"
)

type TransactionType int

const (
	Unknown                          = -1
	TransferIncoming TransactionType = 1 + iota
	TransferOutgoing
	CardPayment
	ATMWithdrawal
	TransferWeb
	TransferMobile
	Charge
)

type Transaction struct {
	Id              string
	Account         string
	TransactionDate int64
	CurrencyDate    int64
	Type            TransactionType
	Amount          float64
	Balance         float64
	Title           string
	Counterpart     string
	Details         string
}

func (t *Transaction) Hash() string {
	transactionStringified := fmt.Sprint(*t)
	fmt.Println(transactionStringified)
	h := md5.New()
	io.WriteString(h, transactionStringified)

	return fmt.Sprintf("%x", h.Sum(nil))
}
