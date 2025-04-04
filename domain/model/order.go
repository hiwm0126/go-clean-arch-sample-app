package model

import (
	"theapp/constants"
	"time"
)

type Order struct {
	OrderNumber     string
	Status          OrderStatus
	ShipmentDueDate string
	OrderTime       time.Time
}

func NewOrder(orderNumber string, status OrderStatus, shipmentDueDate string, orderTime time.Time) *Order {
	return &Order{
		OrderNumber:     orderNumber,
		Status:          status,
		ShipmentDueDate: shipmentDueDate,
		OrderTime:       orderTime,
	}
}

// CanChangeStatus ステータスが変更可能かどうかを返す
func (o *Order) CanChangeStatus() bool {
	switch o.Status {
	case OrderStatusOrdered: // ステータスがOrderedの場合のみキャンセル可能
		return true
	}
	return false
}

// CanChangeStatusDate ステータスが変更可能な日付かどうかを返す
func (o *Order) CanChangeStatusDate(datetime time.Time) bool {
	shipmentDueDate, err := time.Parse(constants.DateFormat, o.ShipmentDueDate)
	if err != nil {
		return false
	}
	return shipmentDueDate.After(datetime)
}

type OrderStatus string

const (
	OrderStatusOrdered  OrderStatus = "ordered"
	OrderStatusCanceled OrderStatus = "cancelled"
	OrderStatusShipped  OrderStatus = "shipped"
)
