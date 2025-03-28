package service

import (
	"errors"
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
)

type OrderValidatingService interface {
	Execute(order *model.Order, itemsInfos map[string]int) error
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

func (s *orderValidatingService) Execute(order *model.Order, itemsInfos map[string]int) error {
	// 出荷可能期間を取得
	sap, err := s.shippingAcceptablePeriodRepo.Get()
	if err != nil {
		return err
	}

	// 出荷可能期間チェック
	if !sap.IsAcceptableDate(order.OrderTime, order.ShipmentDueDate) {
		return err
	}

	// 出荷可能数取得
	sl, err := s.shipmentLimitRepo.GetShipmentLimitByDate(order.ShipmentDueDate)
	if err != nil {
		return err
	}
	asl, err := s.additionalShipmentLimitRepo.GetByShipmentDueDate(order.ShipmentDueDate)
	if err != nil {
		return err
	}
	sl.AdditionalShipmentLimits = asl

	//出荷可能数チェック
	for productNumber, additionalQuantity := range itemsInfos {
		//現在の出荷予定数取得
		currentPlannedShippingQuantity, err := s.orderItemRepo.GetCurrentPlannedShippingQuantity(order.ShipmentDueDate, productNumber)
		if err != nil {
			return err
		}

		if !sl.CanShipping(currentPlannedShippingQuantity, additionalQuantity, order.ShipmentDueDate) {
			return errors.New("出荷可能数を超えています")
		}
	}

	return nil
}
