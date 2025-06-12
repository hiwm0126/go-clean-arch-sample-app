package commandline

import (
	"strconv"
	"theapp/constants"
	"theapp/domain/model"
	"time"
)

// DataConverter データ変換の責任を持つインターフェース
type DataConverter interface {
	StringToInt(str string) int
	StringToTime(str string) time.Time
	StringSliceToShipmentLimitFlags(flagSlice []string) map[model.DayOfWeek]bool
}

// dataConverter データ変換の実装
type dataConverter struct{}

// NewDataConverter データ変換器のコンストラクタ
func NewDataConverter() DataConverter {
	return &dataConverter{}
}

// StringToInt 文字列を整数に変換
func (dc *dataConverter) StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

// StringToTime 文字列を時刻に変換
func (dc *dataConverter) StringToTime(str string) time.Time {
	result, _ := time.Parse(constants.DateTimeFormat, str)
	return result
}

// StringSliceToShipmentLimitFlags 文字列スライスを出荷制限フラグに変換
func (dc *dataConverter) StringSliceToShipmentLimitFlags(flagSlice []string) map[model.DayOfWeek]bool {
	// 一週間は7日間
	if len(flagSlice) != 7 {
		return nil
	}

	shipmentLimitFlags := make(map[model.DayOfWeek]bool, 7)
	for i, flagStr := range flagSlice {
		dayOfWeek := model.DayOfWeek(i)
		flag := dc.StringToInt(flagStr)
		shipmentLimitFlags[dayOfWeek] = flag == 1
	}
	return shipmentLimitFlags
} 