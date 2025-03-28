package model

type OrderItem struct {
	OrderNumber     string
	ProductNumber   string
	ShipmentDueDate string
}

func NewOrderItem(orderNumber string, productNumber string, shippingDueDate string) *OrderItem {
	return &OrderItem{
		OrderNumber:     orderNumber,
		ProductNumber:   productNumber,
		ShipmentDueDate: shippingDueDate,
	}
}
