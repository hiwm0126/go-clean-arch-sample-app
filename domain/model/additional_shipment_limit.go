package model

import (
	"theapp/constants"
	"time"
)

type AdditionalShipmentLimit struct {
	Quantity int
	From     time.Time
	To       time.Time
}

func NewAdditionalShipmentLimit(quantity int, from string, to string) (*AdditionalShipmentLimit, error) {
	fromTime, err := time.Parse(constants.DateFormat, from)
	if err != nil {
		return nil, err
	}
	toTime, err := time.Parse(constants.DateFormat, to)
	if err != nil {
		return nil, err
	}
	return &AdditionalShipmentLimit{
		Quantity: quantity,
		From:     fromTime,
		To:       toTime,
	}, nil
}
