package mediator

import "context"

type MessageMediator interface {
	Register(h Handler)
	Send(ctx context.Context, msg Message) (any, error)
}

type Handler interface {
	Handle(ctx context.Context, msg Message) (any, error)
	GetType() string
}

type Message interface {
	GetHandlerType() string
}
