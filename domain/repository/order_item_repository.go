package repository

import (
	"context"
	"theapp/domain/model"
)

type OrderItemRepository interface {
	Save(ctx context.Context, orderItem *model.OrderItem) error
	FindByOrderNumber(ctx context.Context, orderNumber string) (map[string][]*model.OrderItem, error)
	GetCurrentPlannedShippingQuantity(ctx context.Context, shipmentDueDate string, productNumber string) (int, error)
	DeleteByOrderNumber(ctx context.Context, orderNumber string) error
}
