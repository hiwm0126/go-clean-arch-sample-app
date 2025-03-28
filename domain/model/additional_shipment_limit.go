package model

import (
	"github.com/hiwm0126/internship_27_test/constants"
	"time"
)

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
