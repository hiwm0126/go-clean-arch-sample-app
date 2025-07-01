package handler

type Handler interface {
	CanHandle(interface{}) bool
	Handler(interface{}) error
}
