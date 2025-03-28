package datastore

import (
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
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
func (r *shippingAcceptablePeriodRepository) Save(p *model.ShippingAcceptablePeriod) error {
	r.shippingAcceptablePeriodData = &ShippingAcceptablePeriod{
		Duration: p.Duration,
	}
	return nil
}

// GetDuration 出荷可能期間を取得する
func (r *shippingAcceptablePeriodRepository) Get() (*model.ShippingAcceptablePeriod, error) {
	return r.shippingAcceptablePeriodData.ToModel(), nil
}
