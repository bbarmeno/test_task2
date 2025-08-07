package service

import (
	"acc_balance/pkg/domain/model"
	"acc_balance/pkg/domain/repository"
	"context"
	"errors"
	"fmt"
)

type AccBalanceService struct {
	repo repository.AccBalanceRepository
}

func NewAccBalanceService(repo repository.AccBalanceRepository) (*AccBalanceService, error) {

	if repo == nil {
		return nil, fmt.Errorf("repository is nil")
	}

	svc := &AccBalanceService{
		repo: repo,
	}

	return svc, nil
}

func (svc *AccBalanceService) AddTransaction(
	ctx context.Context,
	transactionId string,
	amount float64,
	state string,
	userId uint64,
) (model.UserAcc, error) {

	userAcc, err := svc.repo.GetUserAccBalance(ctx, userId)
	if err != nil {
		return model.UserAcc{}, err
	}

	if state == "lose" {

		amount = -amount
		if userAcc.Balance-amount < 0 {
			return model.UserAcc{}, errors.New("not enough balance")
		}
	}

	newTransaction := model.NewTransaction(transactionId, amount, state, userId)

	err = svc.repo.AddTransaction(ctx, newTransaction)
	if err != nil {
		return model.UserAcc{}, err
	}

	err = svc.repo.UpdateUserAccBalance(ctx, newTransaction, userId)

	return model.UserAcc{}, nil
}

func (svc *AccBalanceService) GetUserAccBalance(ctx context.Context, userId uint64) (model.UserAcc, error) {

	userAcc, err := svc.repo.GetUserAccBalance(ctx, userId)
	if err != nil {
		return model.UserAcc{}, err
	}

	return userAcc, nil
}
