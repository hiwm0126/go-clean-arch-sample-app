package repository

import "example.com/internship_27_test/domain/model"

type AdditionalShipmentLimitRepository interface {
	Save(shipment *model.AdditionalShipmentLimit) error
	GetByShipmentDueDate(shipmentDueDate string) ([]*model.AdditionalShipmentLimit, error)
}
