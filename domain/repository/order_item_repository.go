package repository

import (
	"context"
	"example.com/internship_27_test/domain/model"
)

type OrderItemRepository interface {
	Save(ctx context.Context, orderItem *model.OrderItem) error
	FindByOrderNumber(ctx context.Context, orderNumber string) (map[string][]*model.OrderItem, error)
	GetCurrentPlannedShippingQuantity(ctx context.Context, shipmentDueDate string, productNumber string) (int, error)
	DeleteByOrderNumber(ctx context.Context, orderNumber string) error
}
