package commandline

import (
	"context"
	"errors"
	"fmt"

	"theapp/interfaces/commandline/cli"
	"theapp/interfaces/commandline/handler"
	"theapp/interfaces/commandline/internal/cmdname"
)

// dispatchFn は1コマンド種別の実行（ペイロード検証・ユースケース呼び出し・CLI出力）
type dispatchFn func(ctx context.Context, cmd cli.ParsedCommand) error

// Dispatcher は ParsedCommand を map ルックアップで CommandHandler に委譲する。
// map をミューテートするメソッドがあるためポインタレシーバを使うが、本体は値型として Router や appDeps が保持する。
type Dispatcher struct {
	handlers map[cmdname.CommandName]dispatchFn
}

// RegisterHandler は CommandHandler を1件登録する。
func (d *Dispatcher) RegisterHandler(h handler.CommandHandler) error {
	if h == nil {
		return errors.New("commandline: RegisterHandler nil handler")
	}
	return d.Register(h.CommandName(), h.Handle)
}

// Register は任意のコマンド名に関数ハンドラを登録する。既に登録済みならエラー。
func (d *Dispatcher) Register(name cmdname.CommandName, fn dispatchFn) error {
	if fn == nil {
		return errors.New("commandline: Register nil handler")
	}
	if _, exists := d.handlers[name]; exists {
		return fmt.Errorf("commandline: handler already registered for %s", name)
	}
	d.handlers[name] = fn
	return nil
}

// Dispatch は cmd.Kind に対応するハンドラを1件だけ実行する。
func (d *Dispatcher) Dispatch(ctx context.Context, cmd cli.ParsedCommand) error {
	fn, ok := d.handlers[cmd.Kind]
	if !ok {
		return fmt.Errorf("commandline: unhandled command %q", cmd.Kind)
	}
	return fn(ctx, cmd)
}
