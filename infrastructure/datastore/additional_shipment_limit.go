package datastore

import (
	"context"
	"example.com/internship_27_test/constants"
	"example.com/internship_27_test/domain/model"
	"example.com/internship_27_test/domain/repository"
	"time"
)

type AdditionalShipment struct {
	Quantity int
	From     time.Time
	To       time.Time
}

func (a *AdditionalShipment) ToModel() *model.AdditionalShipmentLimit {
	return &model.AdditionalShipmentLimit{
		Quantity: a.Quantity,
		From:     a.From,
		To:       a.To,
	}
}

type additionalShipmentLimitRepository struct {
	additionalShipmentData []*AdditionalShipment
}

func NewAdditionalShipmentLimitRepository() repository.AdditionalShipmentLimitRepository {
	return &additionalShipmentLimitRepository{
		additionalShipmentData: make([]*AdditionalShipment, 0),
	}
}

func (r *additionalShipmentLimitRepository) Save(_ context.Context, shipment *model.AdditionalShipmentLimit) error {
	r.additionalShipmentData = append(r.additionalShipmentData, &AdditionalShipment{
		Quantity: shipment.Quantity,
		From:     shipment.From,
		To:       shipment.To,
	})
	return nil
}

func (r *additionalShipmentLimitRepository) GetByShipmentDueDate(_ context.Context, shipmentDueDate string) ([]*model.AdditionalShipmentLimit, error) {
	shipmentDueDateTime, err := time.Parse(constants.DateFormat, shipmentDueDate)
	if err != nil {
		return nil, err
	}

	var result []*model.AdditionalShipmentLimit
	for _, data := range r.additionalShipmentData {
		if (data.From.Before(shipmentDueDateTime) && data.To.After(shipmentDueDateTime)) || // From,Toの期間内　または
			data.From.Equal(shipmentDueDateTime) || // Fromが一致
			data.To.Equal(data.To) { // Toが一致
			result = append(result, data.ToModel())
		}
	}
	return result, nil
}
