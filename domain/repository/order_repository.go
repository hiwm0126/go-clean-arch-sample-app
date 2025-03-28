package repository

import "example.com/internship_27_test/domain/model"

type OrderRepository interface {
	Save(order *model.Order) error
	Find(id int) (*model.Order, error)
	FindByOrderNumber(orderNumber string) (*model.Order, error)
	GetOrdersByShipmentDueDate(shipmentDueDate string) ([]*model.Order, error)
	UpdateStatus(orderNumber string, status model.OrderStatus) error
}
