package repository

import (
	"context"
	"example.com/internship_27_test/domain/model"
)

type ShippingAcceptablePeriodRepository interface {
	Save(ctx context.Context, p *model.ShippingAcceptablePeriod) error
	Get(ctx context.Context) (*model.ShippingAcceptablePeriod, error)
}
