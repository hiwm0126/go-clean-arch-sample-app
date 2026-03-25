package commandline

import (
	"context"
	"fmt"

	"theapp/interfaces/commandline/internal/cmdname"
)

type Router interface {
	Routing(args [][]string) error
}

type router struct {
	dispatcher *Dispatcher
}

// NewRouter アプリケーションルーターを構築する
func NewRouter() Router {
	deps := newAppDeps()
	return &router{dispatcher: deps.Dispatcher}
}

func (r *router) Routing(args [][]string) error {
	paramFactory := NewParamFactory()
	cmds, err := paramFactory.Create(args)
	if err != nil {
		return err
	}

	if err := validateInitDataQueryCount(cmds); err != nil {
		return err
	}

	for _, cmd := range cmds {
		if err := r.dispatcher.Dispatch(context.Background(), cmd); err != nil {
			return err
		}
	}
	return nil
}

// validateInitDataQueryCount は先頭が INIT_DATA のとき、後続コマンド数が NumOfQuery と一致するか検証する
func validateInitDataQueryCount(cmds []ParsedCommand) error {
	if len(cmds) == 0 {
		return nil
	}
	first := cmds[0]
	if first.Kind != cmdname.CommandNameInitData || first.InitData == nil {
		return nil
	}
	want := first.InitData.NumOfQuery
	got := len(cmds) - 1
	if got != want {
		return fmt.Errorf("query count mismatch: INIT_DATA NumOfQuery=%d, got %d commands after init", want, got)
	}
	return nil
}
