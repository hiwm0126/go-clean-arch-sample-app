package datastore

import (
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
	"time"
)

type Order struct {
	ID              int
	OrderNumber     string
	status          string
	ShipmentDueDate string
	OrderTime       time.Time
}

func (o *Order) ToModel() *model.Order {
	return &model.Order{
		OrderNumber:     o.OrderNumber,
		Status:          model.OrderStatus(o.status),
		ShipmentDueDate: o.ShipmentDueDate,
		OrderTime:       o.OrderTime,
	}
}

type orderRepository struct {
	index      int
	orderTable map[string]*Order
}

func NewOrderRepository() repository.OrderRepository {
	return &orderRepository{
		orderTable: make(map[string]*Order),
		index:      1,
	}
}

func (r *orderRepository) Save(order *model.Order) error {
	r.orderTable[order.OrderNumber] = &Order{
		ID:              r.index,
		OrderNumber:     order.OrderNumber,
		status:          string(order.Status),
		ShipmentDueDate: order.ShipmentDueDate,
		OrderTime:       order.OrderTime,
	}
	r.index++
	return nil
}

func (r *orderRepository) Find(id int) (*model.Order, error) {
	for _, v := range r.orderTable {
		if v.ID == id {
			return v.ToModel(), nil
		}
	}
	return nil, nil
}

func (r *orderRepository) FindByOrderNumber(orderNumber string) (*model.Order, error) {
	v, ok := r.orderTable[orderNumber]
	if !ok {
		return nil, nil
	}
	return v.ToModel(), nil
}

func (r *orderRepository) GetOrdersByShipmentDueDate(shipmentDueDate string) ([]*model.Order, error) {
	orders := []*model.Order{}
	for _, v := range r.orderTable {
		if v.ShipmentDueDate == shipmentDueDate {
			orders = append(orders, v.ToModel())
		}
	}
	return orders, nil
}

func (r *orderRepository) UpdateStatus(orderNumber string, status model.OrderStatus) error {
	v, ok := r.orderTable[orderNumber]
	if !ok {
		return nil
	}
	v.status = string(status)
	return nil
}
