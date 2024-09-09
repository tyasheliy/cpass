package mediator

import (
	"context"
	"errors"
	"fmt"
)

type MessageMediatorImpl struct {
	relations map[string]Handler
}

func NewMessageMediator() *MessageMediatorImpl {
	return &MessageMediatorImpl{
		relations: make(map[string]Handler),
	}
}

func (m *MessageMediatorImpl) Register(h Handler) {
	m.relations[h.GetType()] = h
}

func (m *MessageMediatorImpl) Send(ctx context.Context, msg Message) (any, error) {
	handlerType := msg.GetHandlerType()

	handler, ok := m.relations[handlerType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("handler for %s message type not found", handlerType))
	}

	return handler.Handle(ctx, msg)
}
