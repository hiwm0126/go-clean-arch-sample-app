package handler

import (
	"context"
	"errors"
	"theapp/usecase"
)

type orderHandler struct {
	orderUseCase usecase.OrderUseCase
}

func NewOrderHandler(orderUseCase usecase.OrderUseCase) Handler {
	return &orderHandler{
		orderUseCase: orderUseCase,
	}
}

func (h *orderHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.OrderUseCaseReq)
	return ok
}

func (h *orderHandler) Handler(param interface{}) error {
	req, ok := param.(*usecase.OrderUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for OrderUseCaseReq")
	}

	// 注文処理を実行
	_, err := h.orderUseCase.Order(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
