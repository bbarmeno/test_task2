package mysql

import (
	"acc_balance/pkg/domain/model"
	"time"
)

type UserAcc struct {
	Id          uint64    `db:"ID" json:"id"`
	DateCreated time.Time `db:"DATE_CREATED" json:"date_created"`
	Balance     float64   `db:"BALANCE" json:"balance"`
}

func (u UserAcc) ToUserAcc() (model.UserAcc, error) {

	return model.UserAcc{
		Id:          u.Id,
		DateCreated: u.DateCreated,
		Balance:     u.Balance,
	}, nil
}

func ToUserAcc(u model.UserAcc) (UserAcc, error) {

	return UserAcc{
		Id:          u.Id,
		DateCreated: u.DateCreated,
		Balance:     u.Balance,
	}, nil
}
