package mysql

import (
	"acc_balance/pkg/domain/model"
	"acc_balance/pkg/domain/repository"
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
)

var _ repository.AccBalanceRepository = (*AccBalanceRepository)(nil)

type AccBalanceRepository struct {
	sqlx *sqlx.DB
}

func NewAccBalanceRepository(sqlx *sqlx.DB) *AccBalanceRepository {
	return &AccBalanceRepository{sqlx: sqlx}
}

func (repo *AccBalanceRepository) AddTransaction(
	ctx context.Context,
	transaction model.Transaction,
) error {

	query := `INSERT INTO public.transactions(
		"ID", 
		"DATE_CREATED", 
		"AMOUNT", 
		"STATE", 
		"USER_ID"
	)VALUES(
		:ID,
		:DATE_CREATED,
		:AMOUNT,
		:STATE,
		:USER_ID
		)`

	transactionVars, err := ToTransaction(transaction)
	if err != nil {
		return err
	}

	_, err = repo.sqlx.NamedExecContext(ctx, query, transactionVars)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AccBalanceRepository) UpdateUserAccBalance(
	ctx context.Context,
	transaction model.Transaction,
	userId uint64,
) error {

	args := make([]any, 0)

	args = append(args, transaction.Amount, userId)

	tx, err := repo.sqlx.Beginx()
	if err != nil {
		return err
	}

	err = tx.QueryRowx(`UPDATE public.user_account SET "BALANCE" = "BALANCE" + $1 WHERE "ID" = $2`, args...).Err()
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (repo *AccBalanceRepository) GetUserAccBalance(
	ctx context.Context,
	userId uint64,
) (model.UserAcc, error) {

	args := make([]any, 0)

	args = append(args, userId)

	var res UserAcc

	query := `SELECT 
		"ID", 
		"DATE_CREATED", 
		"BALANCE"
		FROM public.user_account
		WHERE
		"ID" = $1`

	if err := repo.sqlx.GetContext(ctx, &res, query, args...); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return model.UserAcc{}, model.ErrNoUserAccById
		}
		return model.UserAcc{}, err
	}

	userAcc, err := res.ToUserAcc()

	return userAcc, err
}
