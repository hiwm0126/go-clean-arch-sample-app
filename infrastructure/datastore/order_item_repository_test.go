package datastore

import (
	"context"
	"testing"

	"theapp/domain/model"
)

func TestDatastoreOrderItem_ToModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   *OrderItem
		want *model.OrderItem
	}{
		{
			name: "注文番号と商品番号が写像される",
			in:   &OrderItem{OrderNumber: "O-9", ProductNumber: "P-9"},
			want: &model.OrderItem{
				OrderNumber:     "O-9",
				ProductNumber:   "P-9",
				ShipmentDueDate: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.ToModel()
			if got.OrderNumber != tt.want.OrderNumber ||
				got.ProductNumber != tt.want.ProductNumber ||
				got.ShipmentDueDate != tt.want.ShipmentDueDate {
				t.Fatalf("ToModel() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestOrderItemRepository_Save_FindByOrderNumber(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name    string
		save    []*model.OrderItem
		find    string
		wantLen int
	}{
		{
			name: "同一注文番号の明細を商品番号ごとに取得できる",
			save: []*model.OrderItem{
				model.NewOrderItem("ORD-1", "P-A", "2025-10-01"),
				model.NewOrderItem("ORD-1", "P-B", "2025-10-01"),
			},
			find:    "ORD-1",
			wantLen: 2,
		},
		{
			name:    "該当なしは空マップ",
			save:    []*model.OrderItem{model.NewOrderItem("X", "P", "2025-10-01")},
			find:    "NOPE",
			wantLen: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewOrderItemRepository()
			for _, it := range tt.save {
				if err := repo.Save(ctx, it); err != nil {
					t.Fatal(err)
				}
			}
			got, err := repo.FindByOrderNumber(ctx, tt.find)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.wantLen {
				t.Fatalf("len(map) = %d, want %d; got %#v", len(got), tt.wantLen, got)
			}
		})
	}
}

func TestOrderItemRepository_GetCurrentPlannedShippingQuantity(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name     string
		save     []*model.OrderItem
		date     string
		product  string
		wantQty  int
	}{
		{
			name: "同一配送日・商品の件数が合算される",
			save: []*model.OrderItem{
				model.NewOrderItem("O1", "PX", "2025-11-01"),
				model.NewOrderItem("O2", "PX", "2025-11-01"),
			},
			date:    "2025-11-01",
			product: "PX",
			wantQty: 2,
		},
		{
			name:    "別商品はカウントされない",
			save:    []*model.OrderItem{model.NewOrderItem("O1", "PX", "2025-11-01")},
			date:    "2025-11-01",
			product: "OTHER",
			wantQty: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewOrderItemRepository()
			for _, it := range tt.save {
				if err := repo.Save(ctx, it); err != nil {
					t.Fatal(err)
				}
			}
			n, err := repo.GetCurrentPlannedShippingQuantity(ctx, tt.date, tt.product)
			if err != nil {
				t.Fatal(err)
			}
			if n != tt.wantQty {
				t.Fatalf("quantity = %d, want %d", n, tt.wantQty)
			}
		})
	}
}

func TestOrderItemRepository_DeleteByOrderNumber(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name      string
		before    []*model.OrderItem
		delete    string
		findAfter string
		wantLeft  int
	}{
		{
			name: "指定注文番号のエントリが削除される",
			before: []*model.OrderItem{
				model.NewOrderItem("DEL-1", "P1", "2025-12-01"),
			},
			delete:    "DEL-1",
			findAfter: "DEL-1",
			wantLeft:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewOrderItemRepository()
			for _, it := range tt.before {
				if err := repo.Save(ctx, it); err != nil {
					t.Fatal(err)
				}
			}
			if err := repo.DeleteByOrderNumber(ctx, tt.delete); err != nil {
				t.Fatal(err)
			}
			got, err := repo.FindByOrderNumber(ctx, tt.findAfter)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.wantLeft {
				t.Fatalf("after delete len = %d, want %d", len(got), tt.wantLeft)
			}
		})
	}
}
