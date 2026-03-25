package handler

import (
	"context"
	"errors"
	"fmt"

	"theapp/constants"
	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type orderCommandHandler struct {
	uc usecase.OrderUseCase
}

// NewOrderCommandHandler ORDER 用ハンドラ
func NewOrderCommandHandler(uc usecase.OrderUseCase) CommandHandler {
	return &orderCommandHandler{uc: uc}
}

func (h *orderCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameOrder
}

func (h *orderCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.Order == nil {
		return errors.New("commandline: missing Order payload")
	}
	res, err := h.uc.Order(ctx, cmd.Order)
	if err != nil {
		return err
	}
	if res.IsError {
		fmt.Printf("%s Ordered %s Error: the number of available shipments has been exceeded.\n",
			res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
	} else {
		fmt.Printf("%s Ordered %s\n", res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
	}
	return nil
}
