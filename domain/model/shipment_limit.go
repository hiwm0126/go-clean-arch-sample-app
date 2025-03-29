package model

import (
	"theapp/constants"
	"time"
)

type ShipmentLimit struct {
	DayOfWeek                DayOfWeek
	Quantity                 int
	AdditionalShipmentLimits []*AdditionalShipmentLimit
}

type DayOfWeek int

const (
	Sunday    DayOfWeek = 0
	Monday    DayOfWeek = 1
	Tuesday   DayOfWeek = 2
	Wednesday DayOfWeek = 3
	Thursday  DayOfWeek = 4
	Friday    DayOfWeek = 5
	Saturday  DayOfWeek = 6
)

func NewShipmentLimit(dayOfWeek DayOfWeek, quantity int) *ShipmentLimit {
	return &ShipmentLimit{
		DayOfWeek: dayOfWeek,
		Quantity:  quantity,
	}
}


// GetShipmentLimitQuantityByDate 指定された出荷予定日の、出荷制限数を取得する
// shippingDueDate: 出荷予定日
// 出荷予定日が出荷制限数の追加期間内であれば、出荷制限の数量を加算する
func (s *ShipmentLimit) GetShipmentLimitQuantityByDate(shippingDueDate string) int {
	limitQuantity := s.Quantity
	shippingDueDateTime, err := time.Parse(constants.DateFormat, shippingDueDate)
	if err != nil {
		return 0
	}

	for _, additionalShipmentLimit := range s.AdditionalShipmentLimits {
		if additionalShipmentLimit.From.Before(shippingDueDateTime) && additionalShipmentLimit.To.After(shippingDueDateTime) ||
			additionalShipmentLimit.From.Equal(shippingDueDateTime) ||
			additionalShipmentLimit.To.Equal(shippingDueDateTime) {
			limitQuantity += additionalShipmentLimit.Quantity
		}
	}
	
	return limitQuantity
}