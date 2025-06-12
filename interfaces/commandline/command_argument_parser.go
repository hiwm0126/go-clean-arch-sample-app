package commandline

import (
	"errors"
	"theapp/usecase"
)

// CommandArgumentParser 個別コマンドの引数パース責任を持つインターフェース
type CommandArgumentParser interface {
	CanHandle(commandName CommandName) bool
	Parse(args [][]string, converter DataConverter) (interface{}, error)
}

// OrderArgumentParser ORDER コマンドの引数パーサー
type OrderArgumentParser struct{}

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

func NewOrderArgumentParser() CommandArgumentParser {
	return &OrderArgumentParser{}
}

func (p *OrderArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameOrder
}

func (p *OrderArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
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
			quantity := converter.StringToInt(itemInfo[ORDER_QUANTITY_INDEX])
			itemInfos[productNumber] = quantity
		}
	}

	return &usecase.OrderUseCaseReq{
		OrderTime:       converter.StringToTime(firstLine[ORDER_TIME_INDEX]),
		OrderNumber:     secondLine[ORDER_NUMBER_INDEX],
		ShipmentDueDate: secondLine[ORDER_SHIPMENT_DATE_INDEX],
		ItemInfos:       itemInfos,
	}, nil
}

// CancelArgumentParser CANCEL コマンドの引数パーサー
type CancelArgumentParser struct{}

// CANCEL コマンドの引数構造定数
const (
	CANCEL_TIME_INDEX           = 1 // キャンセル時刻のインデックス
	CANCEL_ORDER_NUMBER_INDEX   = 0 // 注文番号のインデックス
	CANCEL_REQUIRED_LINES       = 2 // 必要行数
	CANCEL_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	CANCEL_SECOND_LINE_MIN_COLS = 1 // 2行目の最小列数
)

func NewCancelArgumentParser() CommandArgumentParser {
	return &CancelArgumentParser{}
}

func (p *CancelArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameCancel
}

func (p *CancelArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
	if len(args) != CANCEL_REQUIRED_LINES {
		return nil, errors.New("CANCEL command requires exactly 2 lines")
	}

	firstLine := args[0]
	secondLine := args[1]

	if len(firstLine) < CANCEL_FIRST_LINE_MIN_COLS || len(secondLine) < CANCEL_SECOND_LINE_MIN_COLS {
		return nil, errors.New("invalid CANCEL command format")
	}

	return &usecase.CancelUseCaseReq{
		OrderNumber: secondLine[CANCEL_ORDER_NUMBER_INDEX],
		CancelTime:  converter.StringToTime(firstLine[CANCEL_TIME_INDEX]),
	}, nil
}

// ShipArgumentParser SHIP コマンドの引数パーサー
type ShipArgumentParser struct{}

// SHIP コマンドの引数構造定数
const (
	SHIP_REQUEST_TIME_INDEX  = 1 // 出荷要求時刻のインデックス
	SHIP_REQUIRED_LINES      = 1 // 必要行数
	SHIP_FIRST_LINE_MIN_COLS = 2 // 1行目の最小列数
)

func NewShipArgumentParser() CommandArgumentParser {
	return &ShipArgumentParser{}
}

func (p *ShipArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameShip
}

func (p *ShipArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
	if len(args) != SHIP_REQUIRED_LINES {
		return nil, errors.New("SHIP command requires exactly 1 line")
	}

	firstLine := args[0]

	if len(firstLine) < SHIP_FIRST_LINE_MIN_COLS {
		return nil, errors.New("invalid SHIP command format")
	}

	return &usecase.ShippingUseCaseReq{
		ShipmentRequestTime: converter.StringToTime(firstLine[SHIP_REQUEST_TIME_INDEX]),
	}, nil
}

// ChangeArgumentParser CHANGE コマンドの引数パーサー
type ChangeArgumentParser struct{}

// CHANGE コマンドの引数構造定数
const (
	CHANGE_REQUEST_TIME_INDEX   = 1 // 変更要求時刻のインデックス
	CHANGE_ORDER_NUMBER_INDEX   = 0 // 注文番号のインデックス
	CHANGE_REQUEST_DATE_INDEX   = 1 // 変更要求日のインデックス
	CHANGE_REQUIRED_LINES       = 2 // 必要行数
	CHANGE_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	CHANGE_SECOND_LINE_MIN_COLS = 2 // 2行目の最小列数
)

func NewChangeArgumentParser() CommandArgumentParser {
	return &ChangeArgumentParser{}
}

func (p *ChangeArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameChange
}

func (p *ChangeArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
	if len(args) != CHANGE_REQUIRED_LINES {
		return nil, errors.New("CHANGE command requires exactly 2 lines")
	}

	firstLine := args[0]
	secondLine := args[1]

	if len(firstLine) < CHANGE_FIRST_LINE_MIN_COLS || len(secondLine) < CHANGE_SECOND_LINE_MIN_COLS {
		return nil, errors.New("invalid CHANGE command format")
	}

	return &usecase.ChangeUseCaseReq{
		OrderNumber:       secondLine[CHANGE_ORDER_NUMBER_INDEX],
		ChangeRequestDate: secondLine[CHANGE_REQUEST_DATE_INDEX],
		RequestTime:       converter.StringToTime(firstLine[CHANGE_REQUEST_TIME_INDEX]),
	}, nil
}

// ExpandArgumentParser EXPAND コマンドの引数パーサー
type ExpandArgumentParser struct{}

// EXPAND コマンドの引数構造定数
const (
	EXPAND_REQUEST_TIME_INDEX   = 1 // 拡張要求時刻のインデックス
	EXPAND_FROM_INDEX           = 0 // 開始日のインデックス
	EXPAND_TO_INDEX             = 1 // 終了日のインデックス
	EXPAND_QUANTITY_INDEX       = 2 // 数量のインデックス
	EXPAND_REQUIRED_LINES       = 2 // 必要行数
	EXPAND_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	EXPAND_SECOND_LINE_MIN_COLS = 3 // 2行目の最小列数
)

func NewExpandArgumentParser() CommandArgumentParser {
	return &ExpandArgumentParser{}
}

func (p *ExpandArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameExpand
}

func (p *ExpandArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
	if len(args) != EXPAND_REQUIRED_LINES {
		return nil, errors.New("EXPAND command requires exactly 2 lines")
	}

	firstLine := args[0]
	secondLine := args[1]

	if len(firstLine) < EXPAND_FIRST_LINE_MIN_COLS || len(secondLine) < EXPAND_SECOND_LINE_MIN_COLS {
		return nil, errors.New("invalid EXPAND command format")
	}

	return &usecase.ExpandUseCaseReq{
		ExpandRequestTime: converter.StringToTime(firstLine[EXPAND_REQUEST_TIME_INDEX]),
		From:              secondLine[EXPAND_FROM_INDEX],
		To:                secondLine[EXPAND_TO_INDEX],
		Quantity:          converter.StringToInt(secondLine[EXPAND_QUANTITY_INDEX]),
	}, nil
}

// InitDataArgumentParser INIT_DATA コマンドの引数パーサー
type InitDataArgumentParser struct{}

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

func NewInitDataArgumentParser() CommandArgumentParser {
	return &InitDataArgumentParser{}
}

func (p *InitDataArgumentParser) CanHandle(commandName CommandName) bool {
	return commandName == CommandNameInitData
}

func (p *InitDataArgumentParser) Parse(args [][]string, converter DataConverter) (interface{}, error) {
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
		NumOfProduct:             converter.StringToInt(firstLine[INIT_NUM_PRODUCT_INDEX]),
		ShipmentLimitThreshold:   converter.StringToInt(firstLine[INIT_SHIPMENT_LIMIT_THRESHOLD_INDEX]),
		ShipmentAcceptablePeriod: converter.StringToInt(firstLine[INIT_SHIPMENT_ACCEPTABLE_PERIOD_INDEX]),
		ProductNumberList:        secondLine,
		ShipmentLimitFlags:       converter.StringSliceToShipmentLimitFlags(thirdLine),
		NumOfQuery:               converter.StringToInt(fourthLine[INIT_NUM_QUERY_INDEX]),
	}, nil
}
