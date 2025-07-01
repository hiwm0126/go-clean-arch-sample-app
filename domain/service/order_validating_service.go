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
	shipmentLimitRepo            repository.ShipmentLimitRepository
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository
	additionalShipmentLimitRepo  repository.AdditionalShipmentLimitRepository
}

func NewOrderValidatingService(
	orderItemRepo repository.OrderItemRepository,
	shipmentLimitRepo repository.ShipmentLimitRepository,
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository,
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository,
) OrderValidatingService {
	return &orderValidatingService{
		orderItemRepo:                orderItemRepo,
		shipmentLimitRepo:            shipmentLimitRepo,
		shippingAcceptablePeriodRepo: shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo:  additionalShipmentLimitRepo,
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

	// 出荷可能数取得
	shipmentLimit, err := s.shipmentLimitRepo.GetShipmentLimitByDate(ctx, order.ShipmentDueDate)
	if err != nil {
		return err
	}
	// ShipmentDueDateが含まれる期間で有効な、追加出荷可能数を取得
	asl, err := s.additionalShipmentLimitRepo.GetByShipmentDueDate(ctx, order.ShipmentDueDate)
	if err != nil {
		return err
	}
	// 出荷制限数に追加出荷可能数を設定
	shipmentLimit.SetAdditionalShipmentLimits(asl)

	//出荷制限数チェック
	for productNumber, additionalQuantity := range itemInfos {
		//現在の出荷予定数を商品番号(productNumber)に基づいて取得
		currentPlannedShippingQuantity, err := s.orderItemRepo.GetCurrentPlannedShippingQuantity(ctx, order.ShipmentDueDate, productNumber)
		if err != nil {
			return err
		}

		// 指定された日付の出荷制限数を取得
		limitQuantity := shipmentLimit.GetShipmentLimitQuantityByDate(order.ShipmentDueDate)
		// (指定された日付の出荷可能数 - 現在の出荷予定数) < 今回の注文で追加される出荷数
		if (limitQuantity - currentPlannedShippingQuantity) < additionalQuantity {
			return errors.New("available quantity is not enough")
		}
	}

	return nil
}
