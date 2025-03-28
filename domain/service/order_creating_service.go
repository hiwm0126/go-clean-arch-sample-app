package service

import (
	"errors"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
)

type OrderCreatingService interface {
	Execute(order *model.Order, itemsInfos map[string]int) error
}

type orderService struct {
	orderRepo                    repository.OrderRepository
	orderItemRepo                repository.OrderItemRepository
	productRepo                  repository.ProductRepository
	shipmentLimitRepo            repository.ShipmentLimitRepository
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository
	additionalShipmentLimitRepo  repository.AdditionalShipmentLimitRepository
	orderValidatingService       OrderValidatingService
}

func NewOrderCreatingService(
	orderRepo repository.OrderRepository,
	orderItemRepo repository.OrderItemRepository,
	productRepo repository.ProductRepository,
	shipmentLimitRepo repository.ShipmentLimitRepository,
	shippingAcceptablePeriodRepo repository.ShippingAcceptablePeriodRepository,
	additionalShipmentLimitRepo repository.AdditionalShipmentLimitRepository,
	orderValidatingService OrderValidatingService,
) OrderCreatingService {
	return &orderService{
		orderRepo:                    orderRepo,
		orderItemRepo:                orderItemRepo,
		productRepo:                  productRepo,
		shipmentLimitRepo:            shipmentLimitRepo,
		shippingAcceptablePeriodRepo: shippingAcceptablePeriodRepo,
		additionalShipmentLimitRepo:  additionalShipmentLimitRepo,
		orderValidatingService:       orderValidatingService,
	}
}

func (s *orderService) Execute(order *model.Order, itemsInfos map[string]int) error {

	// 出荷可能かチェック
	err := s.orderValidatingService.Execute(order, itemsInfos)
	if err != nil {
		return errors.New("this order is not allowed")
	}

	// 出荷可能な場合、注文情報を保存
	err = s.orderRepo.Save(order)
	if err != nil {
		return err
	}

	// 注文商品情報を保存
	for productNumber, quantity := range itemsInfos {

		// 商品情報を取得
		product, err := s.productRepo.FindByProductNumber(productNumber)
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
			err = s.orderItemRepo.Save(orderItem)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
