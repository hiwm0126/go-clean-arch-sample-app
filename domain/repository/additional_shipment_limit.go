package repository

import (
	"context"
	"theapp/domain/model"
)

type AdditionalShipmentLimitRepository interface {
	Save(ctx context.Context, shipment *model.AdditionalShipmentLimit) error
	GetByShipmentDueDate(ctx context.Context, shipmentDueDate string) ([]*model.AdditionalShipmentLimit, error)
}
