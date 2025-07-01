package handler

import (
	"context"
	"errors"
	"theapp/usecase"
)

type expandHandler struct {
	expandUseCase usecase.ExpandUseCase
}

func NewExpandHandler(expandUseCase usecase.ExpandUseCase) Handler {
	return &expandHandler{
		expandUseCase: expandUseCase,
	}
}

func (h *expandHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.ExpandUseCaseReq)
	return ok
}

func (h *expandHandler) Handler(param interface{}) error {
	req, ok := param.(*usecase.ExpandUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ExpandUseCaseReq")
	}

	// 拡張処理を実行
	_, err := h.expandUseCase.Expand(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
