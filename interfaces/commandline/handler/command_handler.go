package handler

import (
	"context"

	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/internal/cmdname"
)

// CommandHandler は1種類の CLI コマンドを処理する。新コマンド追加時は実装を1ファイル足し、wiring で Register する。
type CommandHandler interface {
	CommandName() cmdname.CommandName
	Handle(ctx context.Context, cmd cli.ParsedCommand) error
}
