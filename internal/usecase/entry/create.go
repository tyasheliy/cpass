package entry

import "github.com/tyasheliy/cpass/internal/entry"

const createType = "create"

type CreateMessage struct {
	entry entry.Entry
	data  any
}

func NewCreateMessage(entry entry.Entry, data any) *CreateMessage {
	return &CreateMessage{entry: entry, data: data}
}

func (m *CreateMessage) GetHandlerType() string {
	return createType
}

type CreateHandler struct {
}
