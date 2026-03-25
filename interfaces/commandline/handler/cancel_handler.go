package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
	"theapp/usecase"
)

type cancelCommandHandler struct {
	cancelUseCase usecase.CancelUseCase
}

func NewCancelCommandHandler(cancelUseCase usecase.CancelUseCase) CommandHandler {
	return &cancelCommandHandler{
		cancelUseCase: cancelUseCase,
	}
}

func (h *cancelCommandHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.CancelUseCaseReq)
	return ok
}

func (h *cancelCommandHandler) Handle(ctx context.Context, param interface{}) error {
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
