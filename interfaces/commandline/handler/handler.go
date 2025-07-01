package handler

type Handler interface {
	CanHandle(interface{}) bool
	Handle(ctx context.Context, param interface{}) error
}
