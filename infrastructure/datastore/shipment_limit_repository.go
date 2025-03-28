package datastore

import (
	"errors"
	"github.com/hiwm0126/internship_27_test/constants"
	"github.com/hiwm0126/internship_27_test/domain/model"
	"github.com/hiwm0126/internship_27_test/domain/repository"
	"time"
)

// ShipmentLimit 出荷可能数マスタ情報
type ShipmentLimit struct {
	DayOfWeek int // 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
	Quantity  int
}

func (s *ShipmentLimit) ToModel() *model.ShipmentLimit {
	return &model.ShipmentLimit{
		DayOfWeek: model.DayOfWeek(s.DayOfWeek),
		Quantity:  s.Quantity,
	}
}

type shipmentLimitRepository struct {
	shipmentLimitTable map[int]*ShipmentLimit
}

func NewShipmentLimitRepository() repository.ShipmentLimitRepository {
	return &shipmentLimitRepository{
		shipmentLimitTable: make(map[int]*ShipmentLimit),
	}
}

// SaveShipmentLimit 出荷可能数マスタ情報を保存する
func (r *shipmentLimitRepository) Save(shipment *model.ShipmentLimit) error {
	data := &ShipmentLimit{
		DayOfWeek: int(shipment.DayOfWeek),
		Quantity:  shipment.Quantity,
	}
	r.shipmentLimitTable[data.DayOfWeek] = data
	return nil
}

// GetShipmentLimitByDate 指定日の出荷可能数マスタ情報を取得する
func (r *shipmentLimitRepository) GetShipmentLimitByDate(date string) (*model.ShipmentLimit, error) {
	datetime, err := time.Parse(constants.DateFormat, date)
	if err != nil {
		return nil, err
	}
	weekday := datetime.Weekday()
	limit, ok := r.shipmentLimitTable[int(weekday)]
	if !ok {
		return nil, errors.New("Shipment limit not found")
	}

	return limit.ToModel(), nil
}
