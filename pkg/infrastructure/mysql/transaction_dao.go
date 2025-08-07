package mysql

import (
	"acc_balance/pkg/domain/model"
	"time"
)

type Transaction struct {
	Id          string    `db:"ID" json:"id"`
	DateCreated time.Time `db:"DATE_CREATED" json:"date_created"`
	Amount      float64   `db:"AMOUNT" json:"amount"`
	State       string    `db:"STATE" json:"state"`
	UserId      uint64    `db:"USER_ID" json:"userId"`
}

func (t Transaction) ToTransaction() (model.Transaction, error) {

	return model.Transaction{
		Id:          t.Id,
		DateCreated: t.DateCreated,
		Amount:      t.Amount,
		State:       t.State,
		UserId:      t.UserId,
	}, nil
}

func ToTransaction(t model.Transaction) (Transaction, error) {

	return Transaction{
		Id:          t.Id,
		DateCreated: t.DateCreated,
		Amount:      t.Amount,
		State:       t.State,
		UserId:      t.UserId,
	}, nil
}
