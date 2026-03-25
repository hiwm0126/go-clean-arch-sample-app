package handler

import "context"

type CommandHandler interface {
	CanHandle(interface{}) bool
	Handle(ctx context.Context, param interface{}) error
}
