package model_test

import (
	"testing"
	"time"

	"theapp/domain/model"
)

func TestNewOrder(t *testing.T) {
	t.Parallel()

	ot := time.Date(2025, 3, 1, 10, 0, 0, 0, time.UTC)
	tests := []struct {
		name            string
		orderNumber     string
		status          model.OrderStatus
		shipmentDueDate string
		orderTime       time.Time
		want            *model.Order
	}{
		{
			name:            "全フィールドが設定される",
			orderNumber:     "O-1",
			status:          model.OrderStatusOrdered,
			shipmentDueDate: "2025-03-15",
			orderTime:       ot,
			want: &model.Order{
				OrderNumber:     "O-1",
				Status:          model.OrderStatusOrdered,
				ShipmentDueDate: "2025-03-15",
				OrderTime:       ot,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := model.NewOrder(tt.orderNumber, tt.status, tt.shipmentDueDate, tt.orderTime)
			if got.OrderNumber != tt.want.OrderNumber ||
				got.Status != tt.want.Status ||
				got.ShipmentDueDate != tt.want.ShipmentDueDate ||
				!got.OrderTime.Equal(tt.want.OrderTime) {
				t.Fatalf("NewOrder() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestOrder_CanChangeStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		status model.OrderStatus
		want   bool
	}{
		{name: "Ordered なら変更可能", status: model.OrderStatusOrdered, want: true},
		{name: "Cancelled なら変更不可", status: model.OrderStatusCanceled, want: false},
		{name: "Shipped なら変更不可", status: model.OrderStatusShipped, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := model.NewOrder("x", tt.status, "2025-01-01", time.Time{})
			if got := o.CanChangeStatus(); got != tt.want {
				t.Fatalf("CanChangeStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrder_CanChangeStatusDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		shipmentDueDate string
		datetime        time.Time
		want            bool
	}{
		{
			name:            "出荷予定日が基準日より後なら true",
			shipmentDueDate: "2025-02-01",
			datetime:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:            true,
		},
		{
			name:            "出荷予定日が基準日以前なら false",
			shipmentDueDate: "2025-01-01",
			datetime:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:            false,
		},
		{
			name:            "出荷予定日の形式が不正なら false",
			shipmentDueDate: "invalid",
			datetime:        time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			want:            false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			o := model.NewOrder("x", model.OrderStatusOrdered, tt.shipmentDueDate, time.Time{})
			if got := o.CanChangeStatusDate(tt.datetime); got != tt.want {
				t.Fatalf("CanChangeStatusDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
