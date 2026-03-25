package model_test

import (
	"testing"
	"time"

	"theapp/domain/model"
)

func TestNewShippingAcceptablePeriod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		duration int
		want     int
	}{
		{name: "日数がそのまま保持される", duration: 5, want: 5},
		{name: "0日でも生成できる", duration: 0, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := model.NewShippingAcceptablePeriod(tt.duration)
			if got.Duration != tt.want {
				t.Fatalf("Duration = %d, want %d", got.Duration, tt.want)
			}
		})
	}
}

func TestShippingAcceptablePeriod_IsAcceptableDate(t *testing.T) {
	t.Parallel()

	// shipmentDueDate + Duration 日の 00:00 UTC が orderTime より後なら true
	baseDue := "2025-01-10"
	mustParse := func(layout, s string) time.Time {
		t.Helper()
		tt, err := time.Parse(layout, s)
		if err != nil {
			t.Fatal(err)
		}
		return tt
	}

	tests := []struct {
		name               string
		duration           int
		orderTime          time.Time
		shipmentDueDateStr string
		want               bool
	}{
		{
			name:               "締切日内（注文時刻が期限より前）",
			duration:           2,
			orderTime:          mustParse("2006-01-02T15:04:05", "2025-01-11T12:00:00"),
			shipmentDueDateStr: baseDue,
			want:               true,
		},
		{
			name:               "締切日ちょうどの注文は不可（After は厳密）",
			duration:           2,
			orderTime:          mustParse("2006-01-02T15:04:05", "2025-01-12T00:00:00"),
			shipmentDueDateStr: baseDue,
			want:               false,
		},
		{
			name:               "出荷予定日の文字列が不正なら false",
			duration:           3,
			orderTime:          mustParse("2006-01-02T15:04:05", "2025-01-01T00:00:00"),
			shipmentDueDateStr: "invalid",
			want:               false,
		},
		{
			name:               "Duration0なら出荷予定日の翌日0時まで",
			duration:           0,
			orderTime:          mustParse("2006-01-02T15:04:05", "2025-01-09T23:59:59"),
			shipmentDueDateStr: baseDue,
			want:               true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := model.NewShippingAcceptablePeriod(tt.duration)
			if got := s.IsAcceptableDate(tt.orderTime, tt.shipmentDueDateStr); got != tt.want {
				t.Fatalf("IsAcceptableDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
