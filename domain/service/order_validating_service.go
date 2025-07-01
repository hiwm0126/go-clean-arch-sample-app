package service

import (
	"context"
	"errors"
	"theapp/domain/model"
	"theapp/domain/repository"
)

type OrderValidatingService interface {
	Execute(ctx context.Context, order *model.Order, itemsInfos map[string]int) error
}

type orderValidatingService struct {
	orderItemRepo                repository.OrderItemRepository
	shipmentLimitProvider        ShipmentLimitProvider
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository
}

func NewOrderValidatingService(
	orderItemRepo repository.OrderItemRepository,
	shipmentLimitProvider ShipmentLimitProvider,
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository,
) OrderValidatingService {
	return &orderValidatingService{
		orderItemRepo:                orderItemRepo,
		shipmentLimitProvider:        shipmentLimitProvider,
		shippingAcceptablePeriodRepo: shippingAcceptablePeriodRepo,
	}
}

func (s *orderValidatingService) Execute(ctx context.Context, order *model.Order, itemInfos map[string]int) error {
	// 出荷可能期間を取得
	sap, err := s.shippingAcceptablePeriodRepo.Get(ctx)
	if err != nil {
		return err
	}

	// 出荷可能期間チェック
	if !sap.IsAcceptableDate(order.OrderTime, order.ShipmentDueDate) {
		return err
	}

	// 指定された日の出荷制限情報を取得
	shipmentLimit, err := s.shipmentLimitProvider.Provide(ctx, order.ShipmentDueDate)
	if err != nil {
		return err
	}

	//出荷制限数チェック
	for productNumber, additionalQuantity := range itemInfos {
		//現在の出荷予定数を商品番号(productNumber)に基づいて取得
		currentPlannedShippingQuantity, err := s.orderItemRepo.GetCurrentPlannedShippingQuantity(ctx, order.ShipmentDueDate, productNumber)
		if err != nil {
			return err
		}

		// 出荷上限数を取得
		limitQuantity := shipmentLimit.GetShipmentLimitQuantity()
		// (出荷上限数 - 現在の出荷予定数) < 今回の注文で追加される出荷数
		if (limitQuantity - currentPlannedShippingQuantity) < additionalQuantity {
			return errors.New("available quantity is not enough")
		}
	}

	return nil
}
