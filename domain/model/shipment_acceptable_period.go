package model

import (
	"theapp/constants"
	"time"
)

type ShippingAcceptablePeriod struct {
	Duration int
}

func NewShippingAcceptablePeriod(duration int) *ShippingAcceptablePeriod {
	return &ShippingAcceptablePeriod{
		Duration: duration,
	}
}

// Durationに指定された日数後までの範囲に収まっているかどうか
func (s *ShippingAcceptablePeriod) IsAcceptableDate(orderTime time.Time, shipmentDueDateStr string) bool {
	shipmentDueDate, err := time.Parse(constants.DateFormat, shipmentDueDateStr)
	if err != nil {
		return false
	}
	return shipmentDueDate.AddDate(0, 0, s.Duration).After(orderTime)
}
