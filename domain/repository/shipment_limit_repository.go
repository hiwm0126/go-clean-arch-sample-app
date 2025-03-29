package repository

import (
	"context"
	"theapp/domain/model"
)

type ShipmentLimitRepository interface {
	Save(ctx context.Context, shipment *model.ShipmentLimit) error
	GetShipmentLimitByDate(ctx context.Context, date string) (*model.ShipmentLimit, error)
}
