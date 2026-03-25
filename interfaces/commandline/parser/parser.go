package parser

import "theapp/interfaces/commandline/internal/cmdname"

// CommandArgumentParser 個別コマンドの引数パース責任を持つインターフェース
type CommandArgumentParser interface {
	CanHandle(commandName cmdname.CommandName) bool
	Parse(args [][]string) (interface{}, error)
}
