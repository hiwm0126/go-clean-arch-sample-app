package model_test

import (
	"testing"
	"time"

	"theapp/domain/model"
)

func TestNewAdditionalShipmentLimit(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		qty     int
		from    string
		to      string
		wantErr bool
		check   func(t *testing.T, got *model.AdditionalShipmentLimit)
	}{
		{
			name:    "正常な日付範囲で生成できる",
			qty:     10,
			from:    "2025-01-01",
			to:      "2025-01-31",
			wantErr: false,
			check: func(t *testing.T, got *model.AdditionalShipmentLimit) {
				t.Helper()
				if got.Quantity != 10 {
					t.Fatalf("Quantity = %d", got.Quantity)
				}
				wantFrom := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
				wantTo := time.Date(2025, 1, 31, 0, 0, 0, 0, time.UTC)
				if !got.From.Equal(wantFrom) || !got.To.Equal(wantTo) {
					t.Fatalf("From=%v To=%v, want From=%v To=%v", got.From, got.To, wantFrom, wantTo)
				}
			},
		},
		{
			name:    "From が不正ならエラー",
			qty:     1,
			from:    "not-a-date",
			to:      "2025-01-31",
			wantErr: true,
		},
		{
			name:    "To が不正ならエラー",
			qty:     1,
			from:    "2025-01-01",
			to:      "bad",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := model.NewAdditionalShipmentLimit(tt.qty, tt.from, tt.to)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tt.check != nil {
				tt.check(t, got)
			}
		})
	}
}
