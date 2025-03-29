package repository

import (
	"context"
	"theapp/domain/model"
)

type ShippingAcceptablePeriodRepository interface {
	Save(ctx context.Context, p *model.ShippingAcceptablePeriod) error
	Get(ctx context.Context) (*model.ShippingAcceptablePeriod, error)
}
