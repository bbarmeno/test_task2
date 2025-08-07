package acc_balance

import (
	"acc_balance/pkg/domain/repository"
	"acc_balance/pkg/domain/service"
	"acc_balance/pkg/infrastructure/mysql"
	"acc_balance/pkg/presentation/http"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AccBalance struct {
	Service *service.AccBalanceService
	Repo    repository.AccBalanceRepository
	Handler *http.AccBalanceHandler
}

func (a AccBalance) Setup(app *fiber.App) error {
	a.Handler.SetupRoutes(app)
	return nil
}

func Setup(
	storage *sqlx.DB,
	app *fiber.App,
) error {
	accBalance, err := GetAccBalance(storage)
	if err != nil {
		return fmt.Errorf("failed to create balance: %w", err)
	}

	if err := accBalance.Setup(app); err != nil {
		return fmt.Errorf("failed to setup balance: %w", err)
	}

	return nil
}

func GetAccBalance(storage *sqlx.DB) (AccBalance, error) {
	var (
		conn = storage
	)

	accBalanceRepo := mysql.NewAccBalanceRepository(conn)
	accBalance, err := service.NewAccBalanceService(accBalanceRepo)
	if err != nil {
		return AccBalance{}, fmt.Errorf("failed to create account balance service: %w", err)
	}

	accBalanceHandler, err := http.NewAccBalanceServHandler(accBalance)
	if err != nil {
		return AccBalance{}, fmt.Errorf("failed to create account balance handler: %w", err)
	}

	return AccBalance{
		Service: accBalance,
		Repo:    accBalanceRepo,
		Handler: accBalanceHandler,
	}, nil
}
