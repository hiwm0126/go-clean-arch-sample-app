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

type cancelCommandHandler struct {
	uc usecase.CancelUseCase
}

// NewCancelCommandHandler CANCEL 用ハンドラ
func NewCancelCommandHandler(uc usecase.CancelUseCase) CommandHandler {
	return &cancelCommandHandler{uc: uc}
}

func (h *cancelCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameCancel
}

func (h *cancelCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.Cancel == nil {
		return errors.New("commandline: missing Cancel payload")
	}
	res, err := h.uc.Cancel(ctx, cmd.Cancel)
	if err != nil {
		return err
	}
	fmt.Printf("%s Canceled %s\n", res.CancelTime.Format(constants.DateTimeFormat), res.OrderNumber)
	return nil
}
