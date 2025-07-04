package handler

import (
	"context"
	"errors"
	"theapp/usecase"
)

type initDataHandler struct {
	dataInitUseCase usecase.DataInitializationUseCase
}

func NewInitDataHandler(dataInitUseCase usecase.DataInitializationUseCase) Handler {
	return &initDataHandler{
		dataInitUseCase: dataInitUseCase,
	}
}

func (h *initDataHandler) CanHandle(param interface{}) bool {
	_, ok := param.(*usecase.DataInitializationUseCaseReq)
	return ok
}

func (h *initDataHandler) Handle(ctx context.Context, param interface{}) error {
	req, ok := param.(*usecase.DataInitializationUseCaseReq)
	if !ok {
		return errors.New("invalid parameter type for DataInitializationUseCaseReq")
	}

	// データ初期化の処理を実行
	err := h.dataInitUseCase.InitData(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
