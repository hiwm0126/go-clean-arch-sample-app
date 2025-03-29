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

func (s *orderValidatingService) Execute(ctx context.Context, order *model.Order, itemsInfos map[string]int) error {
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
	shipmentLimit.AdditionalShipmentLimits = asl

	//出荷可能数チェック
	for productNumber, additionalQuantity := range itemsInfos {
		//現在の出荷予定数取得
		currentPlannedShippingQuantity, err := s.orderItemRepo.GetCurrentPlannedShippingQuantity(ctx, order.ShipmentDueDate, productNumber)
		if err != nil {
			return err
		}

		// 出荷可能数チェック
		if !shipmentLimit.CanShipping(currentPlannedShippingQuantity, additionalQuantity, order.ShipmentDueDate) {
			return errors.New("出荷可能数を超えています")
		}
	}

	return nil
}
