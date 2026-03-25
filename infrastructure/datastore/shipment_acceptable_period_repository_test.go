package datastore

import (
	"context"
	"testing"

	"theapp/domain/model"
	"theapp/domain/repository"
)

func TestShippingAcceptablePeriod_ToModel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   *ShippingAcceptablePeriod
		want int
	}{
		{name: "Durationが引き継がれる", in: &ShippingAcceptablePeriod{Duration: 14}, want: 14},
		{name: "0日", in: &ShippingAcceptablePeriod{Duration: 0}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.in.ToModel()
			if got.Duration != tt.want {
				t.Fatalf("Duration = %d, want %d", got.Duration, tt.want)
			}
		})
	}
}

func TestShippingAcceptablePeriodRepository_Save_and_Get(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name     string
		sequence []*model.ShippingAcceptablePeriod
		want     int
	}{
		{
			name:     "Saveした値がGetで取得できる",
			sequence: []*model.ShippingAcceptablePeriod{model.NewShippingAcceptablePeriod(7)},
			want:     7,
		},
		{
			name: "連続Saveで最後の値が残る",
			sequence: []*model.ShippingAcceptablePeriod{
				model.NewShippingAcceptablePeriod(1),
				model.NewShippingAcceptablePeriod(3),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			repo := NewShippingAcceptablePeriodRepository()
			for _, p := range tt.sequence {
				if err := repo.Save(ctx, p); err != nil {
					t.Fatal(err)
				}
			}
			got, err := repo.Get(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if got.Duration != tt.want {
				t.Fatalf("Get().Duration = %d, want %d", got.Duration, tt.want)
			}
		})
	}
}

func TestShippingAcceptablePeriodRepository_Get_initialState(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	tests := []struct {
		name string
		repo func() repository.ShippingAcceptablePeriodRepository
		want int
	}{
		{
			name: "初期化直後はDuration0",
			repo: NewShippingAcceptablePeriodRepository,
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := tt.repo().Get(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if got.Duration != tt.want {
				t.Fatalf("Duration = %d, want %d", got.Duration, tt.want)
			}
		})
	}
}
