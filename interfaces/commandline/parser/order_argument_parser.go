package parser

import (
	"errors"
	"theapp/application"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type orderArgumentParser struct{}

// ORDER コマンドの引数構造定数
const (
	ORDER_TIME_INDEX           = 1 // 注文時刻のインデックス
	ORDER_NUMBER_INDEX         = 0 // 注文番号のインデックス
	ORDER_SHIPMENT_DATE_INDEX  = 1 // 出荷予定日のインデックス
	ORDER_PRODUCT_NUMBER_INDEX = 0 // 商品番号のインデックス
	ORDER_QUANTITY_INDEX       = 1 // 数量のインデックス
	ORDER_MIN_REQUIRED_LINES   = 3 // 最小必要行数
	ORDER_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	ORDER_SECOND_LINE_MIN_COLS = 2 // 2行目の最小列数
	ORDER_ITEM_MIN_COLS        = 2 // 商品情報の最小列数
)

// NewOrderArgumentParser ORDER コマンドの引数パーサーを返す
func NewOrderArgumentParser() CommandArgumentParser {
	return &orderArgumentParser{}
}

func (p *orderArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameOrder
}

func (p *orderArgumentParser) Parse(args [][]string) (interface{}, error) {
	if len(args) < ORDER_MIN_REQUIRED_LINES {
		return nil, errors.New("ORDER command requires at least 3 lines")
	}

	firstLine := args[0]
	secondLine := args[1]
	itemInfoList := args[2:]

	if len(firstLine) < ORDER_FIRST_LINE_MIN_COLS || len(secondLine) < ORDER_SECOND_LINE_MIN_COLS {
		return nil, errors.New("invalid ORDER command format")
	}

	itemInfos := make(map[string]int)
	for _, itemInfo := range itemInfoList {
		if len(itemInfo) >= ORDER_ITEM_MIN_COLS {
			productNumber := itemInfo[ORDER_PRODUCT_NUMBER_INDEX]
			quantity := application.StringToInt(itemInfo[ORDER_QUANTITY_INDEX])
			itemInfos[productNumber] = quantity
		}
	}

	return &usecase.OrderUseCaseReq{
		OrderTime:       application.StringToTime(firstLine[ORDER_TIME_INDEX]),
		OrderNumber:     secondLine[ORDER_NUMBER_INDEX],
		ShipmentDueDate: secondLine[ORDER_SHIPMENT_DATE_INDEX],
		ItemInfos:       itemInfos,
	}, nil
}
