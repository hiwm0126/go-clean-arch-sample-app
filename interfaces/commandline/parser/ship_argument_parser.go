package parser

import (
	"errors"
	"theapp/application"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type shipArgumentParser struct{}

// SHIP コマンドの引数構造定数
const (
	SHIP_REQUEST_TIME_INDEX  = 1 // 出荷要求時刻のインデックス
	SHIP_REQUIRED_LINES      = 1 // 必要行数
	SHIP_FIRST_LINE_MIN_COLS = 2 // 1行目の最小列数
)

// NewShipArgumentParser SHIP コマンドの引数パーサーを返す
func NewShipArgumentParser() CommandArgumentParser {
	return &shipArgumentParser{}
}

func (p *shipArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameShip
}

func (p *shipArgumentParser) Parse(args [][]string) (interface{}, error) {
	if len(args) != SHIP_REQUIRED_LINES {
		return nil, errors.New("SHIP command requires exactly 1 line")
	}

	firstLine := args[0]

	if len(firstLine) < SHIP_FIRST_LINE_MIN_COLS {
		return nil, errors.New("invalid SHIP command format")
	}

	return &usecase.ShippingUseCaseReq{
		ShipmentRequestTime: application.StringToTime(firstLine[SHIP_REQUEST_TIME_INDEX]),
	}, nil
}
