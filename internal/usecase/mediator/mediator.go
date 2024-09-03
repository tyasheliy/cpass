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

func (m *MessageMediatorImpl) Register(h Handler) error {
	_, exists := m.relations[h.GetType()]
	if exists {
		return errors.New(fmt.Sprintf("%s already registered", h.GetType()))
	}

	m.relations[h.GetType()] = h

	return nil
}

func (m *MessageMediatorImpl) Send(ctx context.Context, msg Message) (any, error) {
	handlerType := msg.GetHandlerType()

	handler, ok := m.relations[handlerType]
	if !ok {
		return nil, errors.New(fmt.Sprintf("handler for %s message type not found", handlerType))
	}

	return handler.Handle(ctx, msg)
}
