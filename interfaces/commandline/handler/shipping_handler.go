package handler

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"strings"
	"theapp/constants"
	"theapp/usecase"
)

type shippingHandler struct {
	shippingUseCase usecase.ShippingUseCase
}

func NewShippingHandler(shippingUseCase usecase.ShippingUseCase) Handler {
	return &shippingHandler{
		shippingUseCase: shippingUseCase,
	}
}

func (h *shippingHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.ShippingUseCaseReq)
	return ok
}

func (h *shippingHandler) Handle(param interface{}) error {
	req, ok := param.(*usecase.ShippingUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ShippingUseCaseReq")
	}

	// 配送処理を実行
	res, err := h.shippingUseCase.Ship(context.Background(), req)
	if err != nil {
		return err
	}

	// 標準出力の生成
	fmt.Printf("%s Shipped %v Orders\n", res.ShipmentRequestTime.Format(constants.DateTimeFormat), len(res.Orders))
	var orderNumbers []string
	for _, order := range res.Orders {
		orderNumbers = append(orderNumbers, order.OrderNumber)
	}
	slices.Sort(orderNumbers)
	fmt.Println(strings.Join(orderNumbers, " "))

	return nil
}
