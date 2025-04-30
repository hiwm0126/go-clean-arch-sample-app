package usecase

import (
	"context"
	"theapp/domain/model"
	"theapp/domain/service"
	"time"
)

type OrderUseCaseReq struct {
	OrderNumber     string
	OrderTime       time.Time
	ShipmentDueDate string
	ItemInfos       map[string]int
}

type OrderUseCaseRes struct {
	OrderTime   time.Time
	OrderNumber string
	IsError     bool
}

type OrderUseCase interface {
	Order(ctx context.Context, req *OrderUseCaseReq) (*OrderUseCaseRes, error)
}

type orderUseCase struct {
	orderFactory           service.OrderFactory
	orderValidatingService service.OrderValidatingService
}

func NewOrderUseCase(
	orderFactory service.OrderFactory,
	orderValidatingService service.OrderValidatingService,
) OrderUseCase {
	return &orderUseCase{
		orderFactory:           orderFactory,
		orderValidatingService: orderValidatingService,
	}
}

func (o *orderUseCase) Order(ctx context.Context, req *OrderUseCaseReq) (*OrderUseCaseRes, error) {

	// 注文情報モデルを作成
	order := model.NewOrder(req.OrderNumber, model.OrderStatusOrdered, req.ShipmentDueDate, req.OrderTime)

	// 出荷可能かチェック
	err := o.orderValidatingService.Execute(ctx, order, req.ItemInfos)
	if err != nil {
		return &OrderUseCaseRes{req.OrderTime, req.OrderNumber, true}, nil
	}

	// 注文を作成
	err = o.orderFactory.Execute(ctx, order, req.ItemInfos)
	if err != nil {
		return &OrderUseCaseRes{req.OrderTime, req.OrderNumber, true}, nil
	}

	return &OrderUseCaseRes{req.OrderTime, order.OrderNumber, false}, nil
}
