package datastore

import (
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
)

type OrderItem struct {
	OrderNumber   string
	ProductNumber string
}

func (i *OrderItem) ToModel() *model.OrderItem {
	return &model.OrderItem{
		OrderNumber:   i.OrderNumber,
		ProductNumber: i.ProductNumber,
	}
}

type orderItemRepository struct {
	orderItemData map[string]map[string]map[string][]*OrderItem // 配送日付(ShippingDueDate),商品番号(ProductNumber),注文番号(OrderNumber)の順でキーが設定される
}

func NewOrderItemRepository() repository.OrderItemRepository {
	return &orderItemRepository{
		orderItemData: make(map[string]map[string]map[string][]*OrderItem),
	}
}

func (r *orderItemRepository) Save(orderItem *model.OrderItem) error {
	orderItemData := &OrderItem{
		OrderNumber:   orderItem.OrderNumber,
		ProductNumber: orderItem.ProductNumber,
	}
	_, ok := r.orderItemData[orderItem.ShipmentDueDate]
	if !ok {
		r.orderItemData[orderItem.ShipmentDueDate] = make(map[string]map[string][]*OrderItem)
	}
	_, ok = r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber]
	if !ok {
		r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber] = make(map[string][]*OrderItem)
	}
	_, ok = r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber][orderItem.OrderNumber]
	if !ok {
		r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber][orderItem.OrderNumber] = make([]*OrderItem, 0)
	}
	r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber][orderItem.OrderNumber] = append(r.orderItemData[orderItem.ShipmentDueDate][orderItem.ProductNumber][orderItem.OrderNumber], orderItemData)
	return nil
}

// FindByOrderNumber productNumberをKeyにして、OrderItemを取得する
func (r *orderItemRepository) FindByOrderNumber(orderNumber string) (map[string][]*model.OrderItem, error) {
	result := make(map[string][]*model.OrderItem)

	for _, productNumbers := range r.orderItemData {
		for productNumber, orderNumbers := range productNumbers {
			if _, ok := orderNumbers[orderNumber]; ok {
				for _, orderItem := range orderNumbers[orderNumber] {
					result[productNumber] = append(result[productNumber], orderItem.ToModel())
				}
			}
		}
	}
	return result, nil
}

// GetCurrentPlannedShippingQuantity 配送日付と商品番号を指定して、現在の出荷予定数を取得する
func (r *orderItemRepository) GetCurrentPlannedShippingQuantity(shipmentDueDate string, productNumber string) (int, error) {
	currentPlannedShippingQuantity := 0
	if _, ok := r.orderItemData[shipmentDueDate]; ok {
		if _, ok := r.orderItemData[shipmentDueDate][productNumber]; ok {
			for _, orderItems := range r.orderItemData[shipmentDueDate][productNumber] {
				currentPlannedShippingQuantity += len(orderItems)
			}
		}
	}
	return currentPlannedShippingQuantity, nil
}

// DeleteByOrderNumber 注文番号を指定して、その注文番号に関連する商品を削除する
func (r *orderItemRepository) DeleteByOrderNumber(orderNumber string) error {
	for _, productNumbers := range r.orderItemData {
		for _, orderNumbers := range productNumbers {
			delete(orderNumbers, orderNumber)
		}
	}
	return nil
}
