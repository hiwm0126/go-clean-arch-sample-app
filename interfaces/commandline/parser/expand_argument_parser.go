package parser

import (
	"errors"
	"theapp/application"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type expandArgumentParser struct{}

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

// NewExpandArgumentParser EXPAND コマンドの引数パーサーを返す
func NewExpandArgumentParser() CommandArgumentParser {
	return &expandArgumentParser{}
}

func (p *expandArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameExpand
}

func (p *expandArgumentParser) Parse(args [][]string) (interface{}, error) {
	if len(args) != EXPAND_REQUIRED_LINES {
		return nil, errors.New("EXPAND command requires exactly 2 lines")
	}

	firstLine := args[0]
	secondLine := args[1]

	if len(firstLine) < EXPAND_FIRST_LINE_MIN_COLS || len(secondLine) < EXPAND_SECOND_LINE_MIN_COLS {
		return nil, errors.New("invalid EXPAND command format")
	}

	return &usecase.ExpandUseCaseReq{
		ExpandRequestTime: application.StringToTime(firstLine[EXPAND_REQUEST_TIME_INDEX]),
		From:              secondLine[EXPAND_FROM_INDEX],
		To:                secondLine[EXPAND_TO_INDEX],
		Quantity:          application.StringToInt(secondLine[EXPAND_QUANTITY_INDEX]),
	}, nil
}
