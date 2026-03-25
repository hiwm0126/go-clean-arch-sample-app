package handler

import (
	"context"
	"errors"

	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
	"theapp/usecase"
)

type initDataCommandHandler struct {
	uc usecase.DataInitializationUseCase
}

// NewInitDataCommandHandler INIT_DATA 用ハンドラ
func NewInitDataCommandHandler(uc usecase.DataInitializationUseCase) CommandHandler {
	return &initDataCommandHandler{uc: uc}
}

func (h *initDataCommandHandler) CommandName() cmdname.CommandName {
	return cmdname.CommandNameInitData
}

func (h *initDataCommandHandler) Handle(ctx context.Context, cmd cli.ParsedCommand) error {
	if cmd.InitData == nil {
		return errors.New("commandline: missing InitData payload")
	}
	return h.uc.InitData(ctx, cmd.InitData)
}
