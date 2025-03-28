package usecase

import (
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/service"
	"time"
)

type OrderUseCaseReq struct {
	OrderNumber     string
	OrderTime       time.Time
	ShipmentDueDate string
	ItemsInfos      map[string]int
}

type OrderUseCaseRes struct {
	OrderTime   time.Time
	OrderNumber string
	IsError     bool
}

type OrderUseCase interface {
	Order(req *OrderUseCaseReq) (*OrderUseCaseRes, error)
}

type orderUseCase struct {
	orderCreatingService service.OrderCreatingService
}

func NewOrderUseCase(
	orderService service.OrderCreatingService,
) OrderUseCase {
	return &orderUseCase{
		orderCreatingService: orderService,
	}
}

func (o *orderUseCase) Order(req *OrderUseCaseReq) (*OrderUseCaseRes, error) {

	// 注文情報モデルを作成
	order := model.NewOrder(req.OrderNumber, model.OrderStatusOrdered, req.ShipmentDueDate, req.OrderTime)

	// 注文を作成
	err := o.orderCreatingService.Execute(order, req.ItemsInfos)
	if err != nil {
		return &OrderUseCaseRes{req.OrderTime, req.OrderNumber, true}, nil
	}

	return &OrderUseCaseRes{req.OrderTime, order.OrderNumber, false}, nil
}
