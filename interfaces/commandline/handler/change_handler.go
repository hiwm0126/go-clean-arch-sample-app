package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
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

func (h *changeHandler) Handle(ctx context.Context, param interface{}) error {
	req, ok := param.(*usecase.ChangeUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ChangeUseCaseReq")
	}

	// 変更処理を実行
	res, err := h.changeUseCase.Change(ctx, req)
	if err != nil {
		return err
	}

	// 標準出力の生成
	if res.IsError {
		fmt.Printf("%s Changed %s Error: the number of available shipments has been exceeded.\n", res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
	} else {
		fmt.Printf("%s Changed %s\n", res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
	}

	return nil
}
