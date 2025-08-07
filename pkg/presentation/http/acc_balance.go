package http

import (
	"acc_balance/pkg/domain/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AccBalanceHandler struct {
	AccBalanceServ service.AccBalanceService
}

func NewAccBalanceServHandler(
	accBalanceServ *service.AccBalanceService,
) (*AccBalanceHandler, error) {

	if accBalanceServ == nil {
		return nil, fmt.Errorf("account balance service is nil")
	}

	return &AccBalanceHandler{
		AccBalanceServ: *accBalanceServ,
	}, nil
}

func (hand *AccBalanceHandler) SetupRoutes(app *fiber.App) {
	app.Get("/user/:userId/balance", hand.GetUserAccBalance)
	app.Post("/user/:userId/transaction", hand.AddTransaction)
}
