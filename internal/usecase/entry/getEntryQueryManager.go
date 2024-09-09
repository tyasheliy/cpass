package entry

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

const getEntryQueryManagerType = "getEntryQueryManager"

type GetEntryQueryManagerMessage struct{}

func NewGetEntryQueryManagerMessage() *GetEntryQueryManagerMessage {
	return &GetEntryQueryManagerMessage{}
}

func (m *GetEntryQueryManagerMessage) GetHandlerType() string {
	return getEntryQueryManagerType
}

type GetEntryQueryManagerHandler struct {
	manager *entry.QueryManager
}

func NewGetEntryQueryManagerHandler(manager *entry.QueryManager) *GetEntryQueryManagerHandler {
	return &GetEntryQueryManagerHandler{manager: manager}
}

func (h *GetEntryQueryManagerHandler) Handle(ctx context.Context, msg mediator.Message) (any, error) {
	return h.manager, nil
}

func (h *GetEntryQueryManagerHandler) GetType() string {
	return getEntryQueryManagerType
}
