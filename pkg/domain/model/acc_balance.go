package model

import (
	"errors"
	"time"
)

var (
	ErrNoUserAccById = errors.New("no user acc with such id")
)

func NewUserAcc() UserAcc {

	userAcc := UserAcc{
		DateCreated: time.Now(),
		Balance:     0.00,
	}

	return userAcc
}

type UserAcc struct {
	Id          uint64
	DateCreated time.Time
	Balance     float64
}
