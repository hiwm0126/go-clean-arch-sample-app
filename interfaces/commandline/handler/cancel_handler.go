package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
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

func (h *cancelHandler) Handle(ctx context.Context, param interface{}) error {
	req, ok := param.(*usecase.CancelUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for CancelUseCaseReq")
	}

	// キャンセル処理を実行
	res, err := h.cancelUseCase.Cancel(ctx, req)
	if err != nil {
		return err
	}

	// 標準出力の生成
	fmt.Printf("%s Canceled %s\n", res.CancelTime.Format(constants.DateTimeFormat), res.OrderNumber)

	return nil
}
