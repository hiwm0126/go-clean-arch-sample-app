package handler

import (
	"context"
	"errors"
	"theapp/usecase"
)

type cancelHandler struct {
	cancelUseCase usecase.CancelUseCase
}

func NewCancelHandler(cancelUseCase usecase.CancelUseCase) Handler {
	return &cancelHandler{
		cancelUseCase: cancelUseCase,
	}
}

func (h *cancelHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.CancelUseCaseReq)
	return ok
}

func (h *cancelHandler) Handler(param interface{}) error {
	req, ok := param.(*usecase.CancelUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for CancelUseCaseReq")
	}

	// キャンセル処理を実行
	_, err := h.cancelUseCase.Cancel(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
