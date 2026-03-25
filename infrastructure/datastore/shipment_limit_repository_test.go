package datastore

import (
	"context"
	"strings"
	"testing"
	"time"

	"theapp/domain/model"
	"theapp/domain/repository"
)

func TestDatastoreShipmentLimit_ToModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   *ShipmentLimit
		want *model.ShipmentLimit
	}{
		{
			name: "曜日と数量がドメインモデルに写像される",
			in:   &ShipmentLimit{DayOfWeek: 3, Quantity: 42},
			want: model.NewShipmentLimit(model.Wednesday, 42),
		},
		{
			name: "日曜は0",
			in:   &ShipmentLimit{DayOfWeek: 0, Quantity: 1},
			want: model.NewShipmentLimit(model.Sunday, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.ToModel()
			if got.DayOfWeek != tt.want.DayOfWeek || got.Quantity != tt.want.Quantity {
				t.Fatalf("ToModel() = (%v,%d), want (%v,%d)", got.DayOfWeek, got.Quantity, tt.want.DayOfWeek, tt.want.Quantity)
			}
		})
	}
}

func TestShipmentLimitRepository_GetShipmentLimitByDate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	// 2025-01-06 は月曜 (weekday=1) — Go の time.Monday と一致
	mondayDate := "2025-01-06"
	if wd := mustParseDate(t, mondayDate).Weekday(); wd != time.Monday {
		t.Fatalf("fixture date must be Monday, got %v", wd)
	}

	tests := []struct {
		name    string
		setup   func(r repository.ShipmentLimitRepository)
		date    string
		wantErr bool
		errSub  string
		check   func(t *testing.T, got *model.ShipmentLimit)
	}{
		{
			name: "登録した曜日の日付で取得できる",
			setup: func(r repository.ShipmentLimitRepository) {
				_ = r.Save(ctx, model.NewShipmentLimit(model.Monday, 99))
			},
			date:    mondayDate,
			wantErr: false,
			check: func(t *testing.T, got *model.ShipmentLimit) {
				t.Helper()
				if got.Quantity != 99 || got.DayOfWeek != model.Monday {
					t.Fatalf("got %+v", got)
				}
			},
		},
		{
			name:    "未登録の曜日はエラー",
			setup:   func(r repository.ShipmentLimitRepository) {},
			date:    mondayDate,
			wantErr: true,
			errSub:  "Shipment limit not found",
		},
		{
			name:    "日付文字列が不正ならパースエラー",
			setup:   func(r repository.ShipmentLimitRepository) {},
			date:    "not-a-date",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewShipmentLimitRepository()
			tt.setup(repo)
			got, err := repo.GetShipmentLimitByDate(ctx, tt.date)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				if tt.errSub != "" && !strings.Contains(err.Error(), tt.errSub) {
					t.Fatalf("error %q should contain %q", err.Error(), tt.errSub)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}

func mustParseDate(t *testing.T, s string) time.Time {
	t.Helper()
	tt, err := time.Parse("2006-01-02", s)
	if err != nil {
		t.Fatal(err)
	}
	return tt
}
