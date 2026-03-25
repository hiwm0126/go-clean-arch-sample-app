package parser

import (
	"errors"
	"theapp/application"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type changeArgumentParser struct{}

// CHANGE コマンドの引数構造定数
const (
	CHANGE_REQUEST_TIME_INDEX   = 1 // 変更要求時刻のインデックス
	CHANGE_ORDER_NUMBER_INDEX   = 0 // 注文番号のインデックス
	CHANGE_REQUEST_DATE_INDEX   = 1 // 変更要求日のインデックス
	CHANGE_REQUIRED_LINES       = 2 // 必要行数
	CHANGE_FIRST_LINE_MIN_COLS  = 2 // 1行目の最小列数
	CHANGE_SECOND_LINE_MIN_COLS = 2 // 2行目の最小列数
)

// NewChangeArgumentParser CHANGE コマンドの引数パーサーを返す
func NewChangeArgumentParser() CommandArgumentParser {
	return &changeArgumentParser{}
}

func (p *changeArgumentParser) CanHandle(commandName cmdname.CommandName) bool {
	return commandName == cmdname.CommandNameChange
}

func (p *changeArgumentParser) Parse(args [][]string) (interface{}, error) {
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
		RequestTime:       application.StringToTime(firstLine[CHANGE_REQUEST_TIME_INDEX]),
	}, nil
}
