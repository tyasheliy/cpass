package entry

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/entry/create"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

const getCreateEntryCommandType = "getCreateEntryCommand"

type GetCreateEntryCommandMessage struct {
	entry entry.Entry
	data  any
}

func (m *GetCreateEntryCommandMessage) GetHandlerType() string {
	return getCreateEntryCommandType
}

type GetCreateEntryCommandHandler struct {
	commandFactory create.CommandFactory
}

func (h *GetCreateEntryCommandHandler) Handle(ctx context.Context, msg mediator.Message) (any, error) {
	typedMsg := msg.(*GetCreateEntryCommandMessage)

	return h.commandFactory.Create(typedMsg.entry, typedMsg.data)
}

func (h *GetCreateEntryCommandHandler) GetType() string {
	return getCreateEntryCommandType
}
