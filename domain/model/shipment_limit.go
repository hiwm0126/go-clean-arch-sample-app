package model

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

// GetShipmentLimitQuantity 出荷制限数を取得する
func (s *ShipmentLimit) GetShipmentLimitQuantity() int {
	// 初期の出荷制限数を設定
	limitQuantity := s.Quantity

	// 出荷制限数を追加する
	for _, additionalLimit := range s.AdditionalShipmentLimits {
		limitQuantity += additionalLimit.Quantity
	}

	return limitQuantity
}

func (s *ShipmentLimit) SetAdditionalShipmentLimits(additionalLimits []*AdditionalShipmentLimit) {
	s.AdditionalShipmentLimits = additionalLimits
}
