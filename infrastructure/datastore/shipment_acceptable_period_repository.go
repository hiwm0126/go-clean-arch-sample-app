package datastore

import (
	"context"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
)

type ShippingAcceptablePeriod struct {
	Duration int
}

func (s *ShippingAcceptablePeriod) ToModel() *model.ShippingAcceptablePeriod {
	return &model.ShippingAcceptablePeriod{
		Duration: s.Duration,
	}
}

type shippingAcceptablePeriodRepository struct {
	shippingAcceptablePeriodData *ShippingAcceptablePeriod
}

func NewShippingAcceptablePeriodRepository() repository.ShippingAcceptablePeriodRepository {
	return &shippingAcceptablePeriodRepository{
		shippingAcceptablePeriodData: &ShippingAcceptablePeriod{},
	}
}

// SaveShippingAcceptablePeriod 出荷可能期間マスタ情報を保存する
func (r *shippingAcceptablePeriodRepository) Save(_ context.Context, p *model.ShippingAcceptablePeriod) error {
	r.shippingAcceptablePeriodData = &ShippingAcceptablePeriod{
		Duration: p.Duration,
	}
	return nil
}

// GetDuration 出荷可能期間を取得する
func (r *shippingAcceptablePeriodRepository) Get(_ context.Context) (*model.ShippingAcceptablePeriod, error) {
	return r.shippingAcceptablePeriodData.ToModel(), nil
}
