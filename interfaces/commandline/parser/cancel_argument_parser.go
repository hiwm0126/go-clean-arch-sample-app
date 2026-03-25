package parser

import (
	"errors"
	"theapp/application"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type cancelArgumentParser struct{}

// CANCEL コマンドの引数構造定数
const (
	CANCEL_TIME_INDEX           = 1 // キャンセル時刻のインデックス
	CANCEL_ORDER_NUMBER_INDEX   = 0 // 注文番号のインデックス
	CANCEL_REQUIRED_LINES       = 2 // 必要行数
	CANCEL_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	CANCEL_SECOND_LINE_MIN_COLS = 1 // 2行目の最小列数
)

// NewCancelArgumentParser CANCEL コマンドの引数パーサーを返す
func NewCancelArgumentParser() CommandArgumentParser {
	return &cancelArgumentParser{}
}

func (p *cancelArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameCancel
}

func (p *cancelArgumentParser) Parse(args [][]string) (interface{}, error) {
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
		CancelTime:  application.StringToTime(firstLine[CANCEL_TIME_INDEX]),
	}, nil
}
