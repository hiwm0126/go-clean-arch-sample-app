package usecase

import (
	"errors"
	"github.com/hiwm0126/internship_27_test/domain/repository"
	"github.com/hiwm0126/internship_27_test/domain/service"
	"time"
)

type CancelUseCaseReq struct {
	OrderNumber string
	CancelTime  time.Time
}

type CancelUseCaseRes struct {
	OrderNumber string
	CancelTime  time.Time
}

type CancelUseCase interface {
	Cancel(req *CancelUseCaseReq) (*CancelUseCaseRes, error)
}

type cancelUseCase struct {
	orderRepo          repository.OrderRepository
	orderCancelService service.OrderCancelService
}

func NewCancelUseCase(
	orderRepo repository.OrderRepository,
	orderCancelService service.OrderCancelService,
) CancelUseCase {
	return &cancelUseCase{
		orderRepo:          orderRepo,
		orderCancelService: orderCancelService,
	}
}

func (c *cancelUseCase) Cancel(req *CancelUseCaseReq) (*CancelUseCaseRes, error) {
	// 注文情報を取得
	order, err := c.orderRepo.FindByOrderNumber(req.OrderNumber)
	if err != nil {
		return nil, err
	}

	// 注文が存在しない場合
	if order == nil {
		return nil, errors.New("order not found")
	}

	// 注文をキャンセル
	err = c.orderCancelService.Execute(order, req.CancelTime)
	if err != nil {
		return nil, err
	}

	return &CancelUseCaseRes{order.OrderNumber, req.CancelTime}, nil
}
