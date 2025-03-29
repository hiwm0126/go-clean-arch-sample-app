package service

import (
	"context"
	"errors"
	"theapp/domain/model"
	"theapp/domain/repository"
)

type OrderFactory interface {
	Execute(ctx context.Context, order *model.Order, itemsInfos map[string]int) error
}

type orderFactory struct {
	orderRepo                    repository.OrderRepository
	orderItemRepo                repository.OrderItemRepository
	productRepo                  repository.ProductRepository
	shipmentLimitRepo            repository.ShipmentLimitRepository
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository
	additionalShipmentLimitRepo  repository.AdditionalShipmentLimitRepository
}

func NewOrderFactory(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
	productRepo repository.ProductRepository,
	shipmentLimitRepo repository.ShipmentLimitRepository,
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository,
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository,
) OrderFactory {
	return &orderFactory{
		orderRepo:                    orderRepo,
		orderItemRepo:                orderItemRepo,
		productRepo:                  productRepo,
		shipmentLimitRepo:            shipmentLimitRepo,
		shippingAcceptablePeriodRepo: shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo:  additionalShipmentLimitRepo,
	}
}

func (s *orderFactory) Execute(ctx context.Context, order *model.Order, itemsInfos map[string]int) error {

	// 注文情報を保存
	err := s.orderRepo.Save(ctx, order)
	if err != nil {
		return err
	}

	// 注文商品情報を保存
	for productNumber, quantity := range itemsInfos {

		// 商品情報を取得
		product, err := s.productRepo.FindByProductNumber(ctx, productNumber)
		if err != nil {
			return err
		}

		// 商品情報が存在しない場合、エラーを返す
		if product == nil {
			return errors.New("product not found")
		}

		// 商品数の数だけ、注文商品情報を保存
		for i := 0; i < quantity; i++ {
			orderItem := model.NewOrderItem(order.OrderNumber, productNumber, order.ShipmentDueDate)
			err = s.orderItemRepo.Save(ctx, orderItem)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
