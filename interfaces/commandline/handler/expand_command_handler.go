package handler

import (
	"context"
	"errors"
	"fmt"

	"theapp/constants"
	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type expandCommandHandler struct {
	uc usecase.ExpandUseCase
}

// NewExpandCommandHandler EXPAND 用ハンドラ
func NewExpandCommandHandler(uc usecase.ExpandUseCase) CommandHandler {
	return &expandCommandHandler{uc: uc}
}

func (h *expandCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameExpand
}

func (h *expandCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.Expand == nil {
		return errors.New("commandline: missing Expand payload")
	}
	res, err := h.uc.Expand(ctx, cmd.Expand)
	if err != nil {
		return err
	}
	fmt.Printf("%s Expanded\n", res.ExpandRequestTime.Format(constants.DateTimeFormat))
	return nil
}
