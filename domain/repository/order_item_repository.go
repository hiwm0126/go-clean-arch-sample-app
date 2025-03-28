package repository

import "github.com/hiwm0126/internship_27_test/domain/model"

type OrderItemRepository interface {
	Save(orderItem *model.OrderItem) error
	FindByOrderNumber(orderNumber string) (map[string][]*model.OrderItem, error)
	GetCurrentPlannedShippingQuantity(shipmentDueDate string, productNumber string) (int, error)
	DeleteByOrderNumber(orderNumber string) error
}
