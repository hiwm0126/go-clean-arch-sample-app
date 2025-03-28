package repository

import (
	"example.com/internship_27_test/domain/model"
)

type ShipmentLimitRepository interface {
	Save(shipment *model.ShipmentLimit) error
	GetShipmentLimitByDate(date string) (*model.ShipmentLimit, error)
}
