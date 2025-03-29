package repository

import (
	"context"
	"example.com/internship_27_test/domain/model"
)

type ShipmentLimitRepository interface {
	Save(ctx context.Context, shipment *model.ShipmentLimit) error
	GetShipmentLimitByDate(ctx context.Context, date string) (*model.ShipmentLimit, error)
}
