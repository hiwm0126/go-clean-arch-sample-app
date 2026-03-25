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

type changeCommandHandler struct {
	uc usecase.ChangeUseCase
}

// NewChangeCommandHandler CHANGE 用ハンドラ
func NewChangeCommandHandler(uc usecase.ChangeUseCase) CommandHandler {
	return &changeCommandHandler{uc: uc}
}

func (h *changeCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameChange
}

func (h *changeCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.Change == nil {
		return errors.New("commandline: missing Change payload")
	}
	res, err := h.uc.Change(ctx, cmd.Change)
	if err != nil {
		return err
	}
	if res.IsError {
		fmt.Printf("%s Changed %s Error: the number of available shipments has been exceeded.\n",
			res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
	} else {
		fmt.Printf("%s Changed %s\n", res.RequestTime.Format(constants.DateTimeFormat), res.OrderNumber)
	}
	return nil
}
