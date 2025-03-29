package usecase

import (
	"context"
	"example.com/internship_27_test/constants"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
	"time"
)

type ShippingUseCaseReq struct {
	ShipmentRequestTime time.Time
}

type ShippingUseCaseRes struct {
	ShipmentRequestTime time.Time
	Orders              []*model.Order
}

type ShippingUseCase interface {
	Ship(ctx context.Context, req *ShippingUseCaseReq) (*ShippingUseCaseRes, error)
}

type shippingUseCase struct {
	orderRepo repository.OrderRepository
}

func NewShippingUseCase(orderRepo repository.OrderRepository) ShippingUseCase {
	return &shippingUseCase{
		orderRepo: orderRepo,
	}
}

func (s *shippingUseCase) Ship(ctx context.Context, req *ShippingUseCaseReq) (*ShippingUseCaseRes, error) {
	// 出荷予定日で、注文を取得
	orders, err := s.orderRepo.GetOrdersByShipmentDueDate(ctx, req.ShipmentRequestTime.Format(constants.DateFormat))
	if err != nil {
		return nil, err
	}

	// 出荷可能な注文をフィルタリング
	var targetOrders []*model.Order
	for _, order := range orders {
		if order.CanChangeStatus() {
			targetOrders = append(targetOrders, order)
		}
	}

	// 配送済みにする
	for _, targetOrder := range targetOrders {
		err = s.orderRepo.UpdateStatus(ctx, targetOrder.OrderNumber, model.OrderStatusShipped)
		if err != nil {
			return nil, err
		}
	}

	return &ShippingUseCaseRes{req.ShipmentRequestTime, targetOrders}, nil
}
