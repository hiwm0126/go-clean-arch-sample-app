package handler

import (
	"context"
	"errors"
	"fmt"
	"theapp/constants"
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

func (h *expandHandler) Handle(ctx context.Context, param interface{}) error {
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
