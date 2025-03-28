package repository

import (
	"github.com/hiwm0126/internship_27_test/domain/model"
)

type ShipmentLimitRepository interface {
	Save(shipment *model.ShipmentLimit) error
	GetShipmentLimitByDate(date string) (*model.ShipmentLimit, error)
}
