package model_test

import (
	"testing"

	"theapp/domain/model"
)

func TestNewOrderItem(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		orderNumber     string
		productNumber   string
		shipmentDueDate string
		want            *model.OrderItem
	}{
		{
			name:            "3フィールドがすべて設定される",
			orderNumber:     "ORD-1",
			productNumber:   "PRD-9",
			shipmentDueDate: "2025-06-01",
			want: &model.OrderItem{
				OrderNumber:     "ORD-1",
				ProductNumber:   "PRD-9",
				ShipmentDueDate: "2025-06-01",
			},
		},
		{
			name:            "空文字でも生成できる",
			orderNumber:     "",
			productNumber:   "",
			shipmentDueDate: "",
			want: &model.OrderItem{
				OrderNumber:     "",
				ProductNumber:   "",
				ShipmentDueDate: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := model.NewOrderItem(tt.orderNumber, tt.productNumber, tt.shipmentDueDate)
			if got.OrderNumber != tt.want.OrderNumber ||
				got.ProductNumber != tt.want.ProductNumber ||
				got.ShipmentDueDate != tt.want.ShipmentDueDate {
				t.Fatalf("NewOrderItem() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
