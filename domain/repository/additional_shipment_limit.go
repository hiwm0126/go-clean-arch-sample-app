package repository

import (
	"context"
	"example.com/internship_27_test/domain/model"
)

type AdditionalShipmentLimitRepository interface {
	Save(ctx context.Context, shipment *model.AdditionalShipmentLimit) error
	GetByShipmentDueDate(ctx context.Context, shipmentDueDate string) ([]*model.AdditionalShipmentLimit, error)
}
