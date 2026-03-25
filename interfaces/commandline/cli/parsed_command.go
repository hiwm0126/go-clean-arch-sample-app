package cli

import (
	"errors"
	"fmt"

	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

// ParsedCommand は ArgumentSeparator と ParamFactory の結果を型付きで表す。
// Kind と、対応するフィールドのうち1つだけが非nilである不変条件を満たす。
type ParsedCommand struct {
	Kind     cmdname.CommandName
	InitData *usecase.DataInitializationUseCaseReq
	Order    *usecase.OrderUseCaseReq
	Cancel   *usecase.CancelUseCaseReq
	Ship     *usecase.ShippingUseCaseReq
	Change   *usecase.ChangeUseCaseReq
	Expand   *usecase.ExpandUseCaseReq
}

// NewParsedCommand はパーサー結果 raw を Kind に応じて ParsedCommand に格納する。
func NewParsedCommand(kind cmdname.CommandName, raw interface{}) (ParsedCommand, error) {
	switch kind {
	case cmdname.CommandNameInitData:
		v, ok := raw.(*usecase.DataInitializationUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, InitData: v}, nil
	case cmdname.CommandNameOrder:
		v, ok := raw.(*usecase.OrderUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, Order: v}, nil
	case cmdname.CommandNameCancel:
		v, ok := raw.(*usecase.CancelUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, Cancel: v}, nil
	case cmdname.CommandNameShip:
		v, ok := raw.(*usecase.ShippingUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, Ship: v}, nil
	case cmdname.CommandNameChange:
		v, ok := raw.(*usecase.ChangeUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, Change: v}, nil
	case cmdname.CommandNameExpand:
		v, ok := raw.(*usecase.ExpandUseCaseReq)
		if !ok {
			return ParsedCommand{}, fmt.Errorf("cli: internal parse type for %s", kind)
		}
		return ParsedCommand{Kind: kind, Expand: v}, nil
	default:
		return ParsedCommand{}, errors.New("cli: unknown command kind")
	}
}
