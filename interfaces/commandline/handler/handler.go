package handler

type Handler interface {
	CanHandle(interface{}) bool
	Handle(interface{}) error
}
