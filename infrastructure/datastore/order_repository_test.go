package datastore

import (
	"context"
	"testing"
	"time"

	"theapp/domain/model"
	"theapp/domain/repository"
)

func TestOrder_ToModel(t *testing.T) {
	t.Parallel()

	ot := time.Date(2025, 4, 1, 9, 0, 0, 0, time.UTC)
	tests := []struct {
		name string
		in   *Order
		want *model.Order
	}{
		{
			name: "永続化用Orderからドメインモデルへ変換できる",
			in: &Order{
				OrderNumber:     "ON-1",
				status:          string(model.OrderStatusOrdered),
				ShipmentDueDate: "2025-04-10",
				OrderTime:       ot,
			},
			want: &model.Order{
				OrderNumber:     "ON-1",
				Status:          model.OrderStatusOrdered,
				ShipmentDueDate: "2025-04-10",
				OrderTime:       ot,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.ToModel()
			if got.OrderNumber != tt.want.OrderNumber ||
				got.Status != tt.want.Status ||
				got.ShipmentDueDate != tt.want.ShipmentDueDate ||
				!got.OrderTime.Equal(tt.want.OrderTime) {
				t.Fatalf("ToModel() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestOrderRepository_operations(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name string
		run  func(t *testing.T, r repository.OrderRepository)
	}{
		{
			name: "Save後にFindでID検索できる",
			run: func(t *testing.T, r repository.OrderRepository) {
				o := model.NewOrder("A1", model.OrderStatusOrdered, "2025-05-01", time.Date(2025, 5, 1, 0, 0, 0, 0, time.UTC))
				if err := r.Save(ctx, o); err != nil {
					t.Fatal(err)
				}
				got, err := r.Find(ctx, 1)
				if err != nil || got == nil || got.OrderNumber != "A1" {
					t.Fatalf("Find(1) = %+v, err=%v", got, err)
				}
			},
		},
		{
			name: "FindByOrderNumberで注文番号検索できる",
			run: func(t *testing.T, r repository.OrderRepository) {
				o := model.NewOrder("B2", model.OrderStatusShipped, "2025-05-02", time.Time{})
				if err := r.Save(ctx, o); err != nil {
					t.Fatal(err)
				}
				got, err := r.FindByOrderNumber(ctx, "B2")
				if err != nil || got == nil || got.Status != model.OrderStatusShipped {
					t.Fatalf("FindByOrderNumber = %+v, err=%v", got, err)
				}
			},
		},
		{
			name: "存在しない注文番号はnil",
			run: func(t *testing.T, r repository.OrderRepository) {
				got, err := r.FindByOrderNumber(ctx, "___none___")
				if err != nil || got != nil {
					t.Fatalf("got %+v err=%v", got, err)
				}
			},
		},
		{
			name: "GetOrdersByShipmentDueDateで同一出荷予定日を集約できる",
			run: func(t *testing.T, r repository.OrderRepository) {
				d := "2025-06-01"
				_ = r.Save(ctx, model.NewOrder("C1", model.OrderStatusOrdered, d, time.Time{}))
				_ = r.Save(ctx, model.NewOrder("C2", model.OrderStatusOrdered, d, time.Time{}))
				_ = r.Save(ctx, model.NewOrder("C3", model.OrderStatusOrdered, "2025-07-01", time.Time{}))
				got, err := r.GetOrdersByShipmentDueDate(ctx, d)
				if err != nil {
					t.Fatal(err)
				}
				if len(got) != 2 {
					t.Fatalf("len = %d, want 2", len(got))
				}
			},
		},
		{
			name: "UpdateStatusでステータスを更新できる",
			run: func(t *testing.T, r repository.OrderRepository) {
				_ = r.Save(ctx, model.NewOrder("D1", model.OrderStatusOrdered, "2025-08-01", time.Time{}))
				if err := r.UpdateStatus(ctx, "D1", model.OrderStatusCanceled); err != nil {
					t.Fatal(err)
				}
				got, err := r.FindByOrderNumber(ctx, "D1")
				if err != nil || got == nil || got.Status != model.OrderStatusCanceled {
					t.Fatalf("after update = %+v err=%v", got, err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewOrderRepository()
			tt.run(t, repo)
		})
	}
}
