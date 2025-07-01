package handler

import (
	"context"
	"errors"
	"theapp/usecase"
)

type changeHandler struct {
	changeUseCase usecase.ChangeUseCase
}

func NewChangeHandler(changeUseCase usecase.ChangeUseCase) Handler {
	return &changeHandler{
		changeUseCase: changeUseCase,
	}
}

func (h *changeHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.ChangeUseCaseReq)
	return ok
}

func (h *changeHandler) Handler(param interface{}) error {
	req, ok := param.(*usecase.ChangeUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ChangeUseCaseReq")
	}

	// 変更処理を実行
	_, err := h.changeUseCase.Change(context.Background(), req)
	if err != nil {
		return err
	}

	return nil
}
