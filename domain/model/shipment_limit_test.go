package model_test

import (
	"testing"

	"theapp/domain/model"
)

func TestNewShipmentLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		dayOfWeek model.DayOfWeek
		quantity  int
	}{
		{name: "月曜と数量を設定", dayOfWeek: model.Monday, quantity: 100},
		{name: "日曜と0件", dayOfWeek: model.Sunday, quantity: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := model.NewShipmentLimit(tt.dayOfWeek, tt.quantity)
			if got.DayOfWeek != tt.dayOfWeek || got.Quantity != tt.quantity {
				t.Fatalf("NewShipmentLimit() = (%v, %d), want (%v, %d)", got.DayOfWeek, got.Quantity, tt.dayOfWeek, tt.quantity)
			}
			if got.AdditionalShipmentLimits != nil {
				t.Fatalf("AdditionalShipmentLimits should be nil, got %#v", got.AdditionalShipmentLimits)
			}
		})
	}
}

func TestShipmentLimit_GetShipmentLimitQuantity(t *testing.T) {
	t.Parallel()

	fromTo := func(t *testing.T, from, to string) *model.AdditionalShipmentLimit {
		t.Helper()
		a, err := model.NewAdditionalShipmentLimit(5, from, to)
		if err != nil {
			t.Fatal(err)
		}
		return a
	}

	tests := []struct {
		name     string
		setup    func(t *testing.T) *model.ShipmentLimit
		wantQty  int
	}{
		{
			name: "追加なしはベース数量のみ",
			setup: func(t *testing.T) *model.ShipmentLimit {
				return model.NewShipmentLimit(model.Tuesday, 20)
			},
			wantQty: 20,
		},
		{
			name: "追加制限の数量が合算される",
			setup: func(t *testing.T) *model.ShipmentLimit {
				s := model.NewShipmentLimit(model.Wednesday, 10)
				s.SetAdditionalShipmentLimits([]*model.AdditionalShipmentLimit{
					fromTo(t, "2025-01-01", "2025-01-02"),
					fromTo(t, "2025-02-01", "2025-02-02"),
				})
				return s
			},
			wantQty: 20,
		},
		{
			name: "追加リストが空ならベースのみ",
			setup: func(t *testing.T) *model.ShipmentLimit {
				s := model.NewShipmentLimit(model.Thursday, 7)
				s.SetAdditionalShipmentLimits([]*model.AdditionalShipmentLimit{})
				return s
			},
			wantQty: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := tt.setup(t)
			if got := s.GetShipmentLimitQuantity(); got != tt.wantQty {
				t.Fatalf("GetShipmentLimitQuantity() = %d, want %d", got, tt.wantQty)
			}
		})
	}
}
