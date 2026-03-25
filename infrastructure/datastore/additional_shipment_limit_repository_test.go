package datastore

import (
	"context"
	"testing"
	"time"

	"theapp/domain/model"
)

func TestAdditionalShipment_ToModel(t *testing.T) {
	t.Parallel()

	from := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name string
		in   *AdditionalShipment
		want *model.AdditionalShipmentLimit
	}{
		{
			name: "数量と期間がドメインモデルへ写像される",
			in:   &AdditionalShipment{Quantity: 8, From: from, To: to},
			want: &model.AdditionalShipmentLimit{Quantity: 8, From: from, To: to},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.ToModel()
			if got.Quantity != tt.want.Quantity || !got.From.Equal(tt.want.From) || !got.To.Equal(tt.want.To) {
				t.Fatalf("ToModel() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestAdditionalShipmentLimitRepository_GetByShipmentDueDate(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	must := func(t *testing.T, q int, from, to string) *model.AdditionalShipmentLimit {
		t.Helper()
		a, err := model.NewAdditionalShipmentLimit(q, from, to)
		if err != nil {
			t.Fatal(err)
		}
		return a
	}

	tests := []struct {
		name       string
		save       []*model.AdditionalShipmentLimit
		queryDate  string
		wantErr    bool
		wantCount  int
		wantSumQty int
	}{
		{
			name:       "出荷予定日が期間内ならヒットする",
			save:       []*model.AdditionalShipmentLimit{must(t, 3, "2025-03-01", "2025-03-31")},
			queryDate:  "2025-03-15",
			wantCount:  1,
			wantSumQty: 3,
		},
		{
			name:       "Fromちょうどの日付も含まれる",
			save:       []*model.AdditionalShipmentLimit{must(t, 1, "2025-04-10", "2025-04-20")},
			queryDate:  "2025-04-10",
			wantCount:  1,
			wantSumQty: 1,
		},
		{
			name:       "Toちょうどの日付も含まれる",
			save:       []*model.AdditionalShipmentLimit{must(t, 2, "2025-04-10", "2025-04-20")},
			queryDate:  "2025-04-20",
			wantCount:  1,
			wantSumQty: 2,
		},
		{
			name:       "期間外は結果に含まれない",
			save:       []*model.AdditionalShipmentLimit{must(t, 5, "2025-05-01", "2025-05-05")},
			queryDate:  "2025-05-10",
			wantCount:  0,
			wantSumQty: 0,
		},
		{
			name:       "日付文字列が不正ならエラー",
			save:       nil,
			queryDate:  "bad-date",
			wantErr:    true,
			wantCount:  0,
			wantSumQty: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewAdditionalShipmentLimitRepository()
			for _, s := range tt.save {
				if err := repo.Save(ctx, s); err != nil {
					t.Fatal(err)
				}
			}
			got, err := repo.GetByShipmentDueDate(ctx, tt.queryDate)
			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != tt.wantCount {
				t.Fatalf("len = %d, want %d", len(got), tt.wantCount)
			}
			sum := 0
			for _, x := range got {
				sum += x.Quantity
			}
			if sum != tt.wantSumQty {
				t.Fatalf("sum Quantity = %d, want %d", sum, tt.wantSumQty)
			}
		})
	}
}
