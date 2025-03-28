package model

import (
	"github.com/hiwm0126/internship_27_test/constants"
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
	shipmentDueDate, _ := time.Parse(constants.DateFormat, shipmentDueDateStr)
	return shipmentDueDate.AddDate(0, 0, s.Duration).After(orderTime)
}
