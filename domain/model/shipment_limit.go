package model

import (
	"github.com/hiwm0126/internship_27_test/constants"
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

// CanShipping 出荷可能かどうかを判定する
// currentPlannedShippingQuantity: 既に計画されている出荷数
// additionalQuantity: 追加で出荷する数量
// shippingDueDate: 出荷予定日
// 出荷予定日が出荷制限数の追加期間内であれば、出荷制限の数量を加算する
func (s *ShipmentLimit) CanShipping(currentPlannedShippingQuantity int, additionalQuantity int, shippingDueDate string) bool {
	limitQuantity := s.Quantity
	shippingDueDateTime, _ := time.Parse(constants.DateFormat, shippingDueDate)
	for _, additionalShipmentLimit := range s.AdditionalShipmentLimits {
		if additionalShipmentLimit.From.Before(shippingDueDateTime) && additionalShipmentLimit.To.After(shippingDueDateTime) ||
			additionalShipmentLimit.From.Equal(shippingDueDateTime) ||
			additionalShipmentLimit.To.Equal(shippingDueDateTime) {
			limitQuantity += additionalShipmentLimit.Quantity
		}
	}
	return currentPlannedShippingQuantity+additionalQuantity <= limitQuantity
}

type AdditionalShipmentLimit struct {
	Quantity int
	From     time.Time
	To       time.Time
}

func NewAdditionalShipmentLimit(quantity int, from string, to string) *AdditionalShipmentLimit {
	fromTime, _ := time.Parse(constants.DateFormat, from)
	toTime, _ := time.Parse(constants.DateFormat, to)
	return &AdditionalShipmentLimit{
		Quantity: quantity,
		From:     fromTime,
		To:       toTime,
	}
}
