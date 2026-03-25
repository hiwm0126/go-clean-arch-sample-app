package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
	"theapp/usecase"
)

type expandCommandHandler struct {
	expandUseCase usecase.ExpandUseCase
}

func NewExpandCommandHandler(expandUseCase usecase.ExpandUseCase) CommandHandler {
	return &expandCommandHandler{
		expandUseCase: expandUseCase,
	}
}

func (h *expandCommandHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.ExpandUseCaseReq)
	return ok
}

func (h *expandCommandHandler) Handle(ctx context.Context, param interface{}) error {
	req, ok := param.(*usecase.ExpandUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for ExpandUseCaseReq")
	}

	// 拡張処理を実行
	res, err := h.expandUseCase.Expand(ctx, req)
	if err != nil {
		return err
	}

	// 標準出力の生成
	fmt.Printf("%s Expanded\n", res.ExpandRequestTime.Format(constants.DateTimeFormat))

	return nil
}
