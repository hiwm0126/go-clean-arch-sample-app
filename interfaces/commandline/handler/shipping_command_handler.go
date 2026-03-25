package handler

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"

	"theapp/constants"
	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type shippingCommandHandler struct {
	uc usecase.ShippingUseCase
}

// NewShippingCommandHandler SHIP 用ハンドラ
func NewShippingCommandHandler(uc usecase.ShippingUseCase) CommandHandler {
	return &shippingCommandHandler{uc: uc}
}

func (h *shippingCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameShip
}

func (h *shippingCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.Ship == nil {
		return errors.New("commandline: missing Ship payload")
	}
	res, err := h.uc.Ship(ctx, cmd.Ship)
	if err != nil {
		return err
	}
	fmt.Printf("%s Shipped %v Orders\n", res.ShipmentRequestTime.Format(constants.DateTimeFormat), len(res.Orders))
	var orderNumbers []string
	for _, o := range res.Orders {
		orderNumbers = append(orderNumbers, o.OrderNumber)
	}
	slices.Sort(orderNumbers)
	fmt.Println(strings.Join(orderNumbers, " "))
	return nil
}
