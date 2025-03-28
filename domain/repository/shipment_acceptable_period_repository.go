package repository

import "github.com/hiwm0126/internship_27_test/domain/model"

type ShippingAcceptablePeriodRepository interface {
	Save(p *model.ShippingAcceptablePeriod) error
	Get() (*model.ShippingAcceptablePeriod, error)
}
