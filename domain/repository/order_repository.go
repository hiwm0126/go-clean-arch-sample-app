package repository

import (
	"context"
	"theapp/domain/model"
)

type OrderRepository interface {
	Save(ctx context.Context, order *model.Order) error
	Find(ctx context.Context, id int) (*model.Order, error)
	FindByOrderNumber(ctx context.Context, orderNumber string) (*model.Order, error)
	GetOrdersByShipmentDueDate(ctx context.Context, shipmentDueDate string) ([]*model.Order, error)
	UpdateStatus(ctx context.Context, orderNumber string, status model.OrderStatus) error
}
