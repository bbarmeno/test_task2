package repository

import (
	"acc_balance/pkg/domain/model"
	"context"
)

type AccBalanceRepository interface {
	AddTransaction(
		ctx context.Context,
		transaction model.Transaction,
	) error

	UpdateUserAccBalance(
		ctx context.Context,
		transaction model.Transaction,
		userId uint64,
	) error

	GetUserAccBalance(
		ctx context.Context,
		userId uint64,
	) (model.UserAcc, error)
}

type GetUserAccBalanceParams struct {
	Id *uint64 `json:"id" query:"id"`
}

func NewDefaultGetUserAccBalanceParams() GetUserAccBalanceParams {
	return GetUserAccBalanceParams{}
}
