package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
	"theapp/usecase"
)

type orderCommandHandler struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderCommandHandler(orderUseCase usecase.OrderUseCase) CommandHandler {
	return &orderCommandHandler{
		orderUseCase: orderUseCase,
	}
}

func (h *orderCommandHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.OrderUseCaseReq)
	return ok
}

func (h *orderCommandHandler) Handle(ctx context.Context, param interface{}) error {
	req, ok := param.(*usecase.OrderUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for OrderUseCaseReq")
	}

	// 注文処理を実行
	res, err := h.orderUseCase.Order(ctx, req)
	if err != nil {
		return err
	}

	// 標準出力の生成
	if res.IsError {
		fmt.Printf("%s Ordered %s Error: the number of available shipments has been exceeded.\n", res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
	} else {
		fmt.Printf("%s Ordered %s\n", res.OrderTime.Format(constants.DateTimeFormat), res.OrderNumber)
	}

	return nil
}
