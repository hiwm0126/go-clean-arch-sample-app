package datastore

import (
	"context"
	"theapp/constants"
	"theapp/domain/model"
	"theapp/domain/repository"
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
    // 出荷予定日をパース
    shipmentDueDateTime, err := time.Parse(constants.DateFormat, shipmentDueDate)
    if err != nil {
        return nil, err
    }

    // 出荷制限データをフィルタリング
    var result []*model.AdditionalShipmentLimit
    for _, data := range r.additionalShipmentData {
        if isWithinRange(data, shipmentDueDateTime) {
            result = append(result, data.ToModel())
        }
    }
    return result, nil
}

// isWithinRange 出荷予定日が追加制限の期間内かを判定する
func isWithinRange(data *AdditionalShipment, date time.Time) bool {
    return (data.From.Before(date) && data.To.After(date)) || // From,Toの期間内
        data.From.Equal(date) || // Fromが一致
        data.To.Equal(date)      // Toが一致
}
