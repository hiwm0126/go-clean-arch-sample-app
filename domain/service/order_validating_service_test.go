package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"theapp/domain/model"
	"theapp/domain/repository"
)

type stubShippingPeriodRepo struct {
	sap *model.ShippingAcceptablePeriod
	err error
}

func (s *stubShippingPeriodRepo) Save(_ context.Context, _ *model.ShippingAcceptablePeriod) error {
	return nil
}

func (s *stubShippingPeriodRepo) Get(_ context.Context) (*model.ShippingAcceptablePeriod, error) {
	return s.sap, s.err
}

type stubShipmentLimitProvider struct {
	limit *model.ShipmentLimit
	err   error
}

func (s *stubShipmentLimitProvider) Provide(_ context.Context, _ string) (*model.ShipmentLimit, error) {
	return s.limit, s.err
}

type stubOrderItemRepo struct {
	qty int
	err error
}

func (s *stubOrderItemRepo) Save(_ context.Context, _ *model.OrderItem) error { return nil }
func (s *stubOrderItemRepo) FindByOrderNumber(_ context.Context, _ string) (map[string][]*model.OrderItem, error) {
	return nil, nil
}
func (s *stubOrderItemRepo) GetCurrentPlannedShippingQuantity(_ context.Context, _, _ string) (int, error) {
	return s.qty, s.err
}
func (s *stubOrderItemRepo) DeleteByOrderNumber(_ context.Context, _ string) error { return nil }

var (
	_ repository.ShippingAcceptablePeriodRepository = (*stubShippingPeriodRepo)(nil)
	_ ShipmentLimitProvider                         = (*stubShipmentLimitProvider)(nil)
	_ repository.OrderItemRepository                = (*stubOrderItemRepo)(nil)
)

func TestOrderValidatingService_Execute(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name    string
		sap     *model.ShippingAcceptablePeriod
		sapErr  error
		order   *model.Order
		items   map[string]int
		limit   *model.ShipmentLimit
		limErr  error
		itemQty int
		itemErr error
		wantErr string
	}{
		{
			name:   "出荷可能期間外なら明示的なエラー",
			sap:    model.NewShippingAcceptablePeriod(0),
			order:  model.NewOrder("O1", model.OrderStatusOrdered, "2025-01-01", time.Date(2025, 1, 10, 0, 0, 0, 0, time.UTC)),
			items:  map[string]int{"P1": 1},
			limit:  model.NewShipmentLimit(model.Wednesday, 100),
			wantErr: "shipment due date is outside acceptable period",
		},
		{
			name:    "出荷可能期間リポジトリエラーは伝播",
			sapErr:  errors.New("db down"),
			order:   model.NewOrder("O1", model.OrderStatusOrdered, "2025-06-15", time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)),
			items:   map[string]int{"P1": 1},
			wantErr: "db down",
		},
		{
			name:  "期間内かつ在庫十分なら成功",
			sap:   model.NewShippingAcceptablePeriod(30),
			order: model.NewOrder("O1", model.OrderStatusOrdered, "2025-06-15", time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)),
			items: map[string]int{"P1": 2},
			limit: model.NewShipmentLimit(model.Sunday, 10),
			itemQty: 0,
		},
		{
			name:  "在庫不足なら available quantity",
			sap:   model.NewShippingAcceptablePeriod(30),
			order: model.NewOrder("O1", model.OrderStatusOrdered, "2025-06-15", time.Date(2025, 6, 1, 12, 0, 0, 0, time.UTC)),
			items: map[string]int{"P1": 5},
			limit: model.NewShipmentLimit(model.Sunday, 3),
			itemQty: 0,
			wantErr: "available quantity is not enough",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			svc := NewOrderValidatingService(
				&stubOrderItemRepo{qty: tt.itemQty, err: tt.itemErr},
				&stubShipmentLimitProvider{limit: tt.limit, err: tt.limErr},
				&stubShippingPeriodRepo{sap: tt.sap, err: tt.sapErr},
			)
			err := svc.Execute(ctx, tt.order, tt.items)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected err: %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("err = %v, want substring %q", err, tt.wantErr)
			}
		})
	}
}
