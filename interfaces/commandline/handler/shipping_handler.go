package handler

import (
	"context"
	"errors"
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

func (h *shippingHandler) Handler(param interface{}) error {
	req, ok := param.(*usecase.ShippingUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ShippingUseCaseReq")
	}

	// 配送処理を実行
	_, err := h.shippingUseCase.Ship(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
