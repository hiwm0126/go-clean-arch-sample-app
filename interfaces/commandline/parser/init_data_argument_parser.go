package parser

import (
	"errors"
	"theapp/application"
	"theapp/domain/model"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type initDataArgumentParser struct{}

// INIT_DATA コマンドの引数構造定数
const (
	INIT_NUM_PRODUCT_INDEX                = 0 // 商品数のインデックス
	INIT_SHIPMENT_LIMIT_THRESHOLD_INDEX   = 1 // 出荷制限閾値のインデックス
	INIT_SHIPMENT_ACCEPTABLE_PERIOD_INDEX = 2 // 出荷可能期間のインデックス
	INIT_NUM_QUERY_INDEX                  = 0 // クエリ数のインデックス
	INIT_REQUIRED_LINES                   = 4 // 必要行数
	INIT_FIRST_LINE_MIN_COLS              = 3 // 1行目の最小列数
	INIT_FOURTH_LINE_MIN_COLS             = 1 // 4行目の最小列数
)

// NewInitDataArgumentParser INIT_DATA コマンドの引数パーサーを返す
func NewInitDataArgumentParser() CommandArgumentParser {
	return &initDataArgumentParser{}
}

func (p *initDataArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameInitData
}

func (p *initDataArgumentParser) Parse(args [][]string) (interface{}, error) {
	if len(args) != INIT_REQUIRED_LINES {
		return nil, errors.New("INIT_DATA command requires exactly 4 lines")
	}

	firstLine := args[0]
	secondLine := args[1]
	thirdLine := args[2]
	fourthLine := args[3]

	if len(firstLine) < INIT_FIRST_LINE_MIN_COLS || len(fourthLine) < INIT_FOURTH_LINE_MIN_COLS {
		return nil, errors.New("invalid INIT_DATA command format")
	}

	return &usecase.DataInitializationUseCaseReq{
		NumOfProduct:             application.StringToInt(firstLine[INIT_NUM_PRODUCT_INDEX]),
		ShipmentLimitThreshold:   application.StringToInt(firstLine[INIT_SHIPMENT_LIMIT_THRESHOLD_INDEX]),
		ShipmentAcceptablePeriod: application.StringToInt(firstLine[INIT_SHIPMENT_ACCEPTABLE_PERIOD_INDEX]),
		ProductNumberList:        secondLine,
		ShipmentLimitFlags:       p.stringSliceToShipmentLimitFlags(thirdLine),
		NumOfQuery:               application.StringToInt(fourthLine[INIT_NUM_QUERY_INDEX]),
	}, nil
}

// stringSliceToShipmentLimitFlags は出荷制限行（7 要素）を曜日ごとの出荷制限フラグへ変換する。要素数が 7 でない場合は nil。
func (p *initDataArgumentParser) stringSliceToShipmentLimitFlags(flagSlice []string) map[model.DayOfWeek]bool {
	if len(flagSlice) != 7 {
		return nil
	}

	out := make(map[model.DayOfWeek]bool, 7)
	for i, flagStr := range flagSlice {
		day := model.DayOfWeek(i)
		out[day] = application.StringToInt(flagStr) == 1
	}
	return out
}
