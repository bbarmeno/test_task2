package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (hand *AccBalanceHandler) AddTransaction(c *fiber.Ctx) error {
	req, err := hand.NewCreateTransactionRequest(c)
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		return fmt.Errorf("error parsing request: %w", err)
	}

	article, err := hand.AccBalanceServ.AddTransaction(
		c.Context(), req.TransactionId, amount, req.State, req.UserId)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(article)
}

type CreateTransactionRequest struct {
	TransactionId string `json:"transactionId" query:"transactionId" form:"transactionId"`
	Amount        string `json:"amount" query:"amount" form:"amount"`
	State         string `json:"state" query:"state" form:"state"`
	UserId        uint64 `json:"userId" query:"userId" form:"userId"`
}

func (hand *AccBalanceHandler) NewCreateTransactionRequest(c *fiber.Ctx) (CreateTransactionRequest, error) {
	aId := c.Params("userId")
	uId, err := strconv.ParseUint(aId, 10, 64)
	if err != nil {
		return CreateTransactionRequest{}, fmt.Errorf("error parsing request: %w", err)
	}

	var req CreateTransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return CreateTransactionRequest{}, err
	}

	if err := req.Validate(); err != nil {
		return CreateTransactionRequest{}, err
	}

	req.UserId = uId

	return req, nil
}

func (req CreateTransactionRequest) Validate() error {

	if req.Amount == "" {
		return errors.New("transaction amount is required")
	}

	if req.TransactionId == "" {
		return errors.New("transaction id is required")
	}

	if req.State == "" {
		return errors.New("transaction state is required")
	}

	return nil
}
