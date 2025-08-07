package model

import (
	"errors"
	"time"
)

var (
	ErrNoTransactionById = errors.New("no transaction with such id")
)

func NewTransaction(transactionId string, amount float64, state string, userId uint64) Transaction {

	transaction := Transaction{
		Id:          transactionId,
		DateCreated: time.Now(),
		Amount:      amount,
		State:       state,
		UserId:      userId,
	}

	return transaction
}

type Transaction struct {
	Id          string
	DateCreated time.Time
	Amount      float64
	State       string
	UserId      uint64
}
