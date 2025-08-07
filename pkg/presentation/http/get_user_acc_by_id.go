package http

import (
	"acc_balance/pkg/domain/model"
	"acc_balance/pkg/domain/repository"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (hand *AccBalanceHandler) GetUserAccBalance(c *fiber.Ctx) error {
	req, err := hand.NewGetUserAccBalanceRequest(c)
	if err != nil {
		return err
	}

	userAcc, err := hand.AccBalanceServ.GetUserAccBalance(c.Context(), *req.GetParams().Id)
	if err != nil {
		return err
	}

	return c.JSON(NewGetUserAccBalanceResponce(userAcc))
}

type GetUserAccBalanceRequest repository.GetUserAccBalanceParams

func (hand *AccBalanceHandler) NewGetUserAccBalanceRequest(c *fiber.Ctx) (GetUserAccBalanceRequest, error) {

	aId := c.Params("userId")
	params := repository.NewDefaultGetUserAccBalanceParams()
	uId, err := strconv.ParseUint(aId, 10, 64)
	if err != nil {
		return GetUserAccBalanceRequest{}, fmt.Errorf("error parsing request: %w", err)
	}
	params.Id = &uId

	return GetUserAccBalanceRequest(params), nil
}

func (req GetUserAccBalanceRequest) GetParams() repository.GetUserAccBalanceParams {
	return repository.GetUserAccBalanceParams(req)
}

type GetUserAccBalanceResponceItem struct {
	Id      uint64 `json:"userId"`
	Balance string `json:"balance"`
}

func NewGetUserAccBalanceResponce(
	userAcc model.UserAcc,
) GetUserAccBalanceResponceItem {

	u := GetUserAccBalanceResponceItem{
		Id:      userAcc.Id,
		Balance: strconv.FormatFloat(userAcc.Balance, 'f', 2, 64),
	}

	return u
}
