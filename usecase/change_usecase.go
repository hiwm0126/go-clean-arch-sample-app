package usecase

import (
	"context"
	"errors"
	"theapp/domain/model"
	"theapp/domain/repository"
	"theapp/domain/service"
	"time"
)

type ChangeUseCaseReq struct {
	OrderNumber       string
	RequestTime       time.Time
	ChangeRequestDate string
}

type ChangeUseCaseRes struct {
	OrderNumber string
	RequestTime time.Time
	IsError     bool
}

type ChangeUseCase interface {
	Change(ctx context.Context, req *ChangeUseCaseReq) (*ChangeUseCaseRes, error)
}

type changeUseCase struct {
	orderRepo              repository.OrderRepository
	orderItemRepo          repository.OrderItemRepository
	orderFactory           service.OrderFactory
	orderCancelService     service.OrderCancelService
	orderValidatingService service.OrderValidatingService
}

func NewChangeUseCase(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
	orderFactory service.OrderFactory,
	orderCancelService service.OrderCancelService,
	orderValidatingService service.OrderValidatingService,
) ChangeUseCase {
	return &changeUseCase{
		orderRepo:              orderRepo,
		orderItemRepo:          orderItemRepo,
		orderFactory:           orderFactory,
		orderCancelService:     orderCancelService,
		orderValidatingService: orderValidatingService,
	}
}

func (c *changeUseCase) Change(ctx context.Context, req *ChangeUseCaseReq) (*ChangeUseCaseRes, error) {

	// 対象の注文情報を取得
	targetOrder, err := c.orderRepo.FindByOrderNumber(ctx, req.OrderNumber)
	if err != nil {
		return &ChangeUseCaseRes{targetOrder.OrderNumber, req.RequestTime, true}, nil
	}

	// 注文が存在しない場合
	if targetOrder == nil {
		return nil, errors.New("targetOrder not found")
	}

	// 注文情報モデルを作成
	newOrder := model.NewOrder(req.OrderNumber, model.OrderStatusOrdered, req.ChangeRequestDate, targetOrder.OrderTime)
	// 注文商品情報を取得
	orderProductItems, err := c.orderItemRepo.FindByOrderNumber(ctx, targetOrder.OrderNumber)
	if err != nil {
		return &ChangeUseCaseRes{targetOrder.OrderNumber, req.RequestTime, true}, nil
	}
	itemsInfos := make(map[string]int)
	for productNumber, items := range orderProductItems {
		itemsInfos[productNumber] = len(items)
	}

	// 新しい出荷日への変更が妥当かチェック
	err = c.orderValidatingService.Execute(ctx, newOrder, itemsInfos)
	if err != nil {
		return &ChangeUseCaseRes{newOrder.OrderNumber, req.RequestTime, true}, nil
	}

	// 元の注文のキャンセル処理
	err = c.orderCancelService.Execute(ctx, targetOrder, req.RequestTime)
	if err != nil {
		return &ChangeUseCaseRes{targetOrder.OrderNumber, req.RequestTime, true}, nil
	}

	// 注文を作成
	err = c.orderFactory.Execute(ctx, newOrder, itemsInfos)
	if err != nil {
		return &ChangeUseCaseRes{newOrder.OrderNumber, req.RequestTime, true}, nil
	}

	return &ChangeUseCaseRes{newOrder.OrderNumber, req.RequestTime, false}, nil
}
