package entry

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entity"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

const getEntriesByTypeType = "GetEntriesByType"

type GetEntriesByTypeMessage struct {
	store     *entity.Store
	entryType entity.EntryType
}

func NewGetEntriesByTypeMessage(store *entity.Store, entryType entity.EntryType) *GetEntriesByTypeMessage {
	return &GetEntriesByTypeMessage{store: store, entryType: entryType}
}

func (m *GetEntriesByTypeMessage) GetHandlerType() string {
	return getEntriesByTypeType
}

type GetEntriesByTypeHandler struct {
	repo entity.EntryRepository
}

func NewGetEntriesByTypeHandler(repo entity.EntryRepository) *GetEntriesByTypeHandler {
	return &GetEntriesByTypeHandler{repo: repo}
}

func (h *GetEntriesByTypeHandler) Handle(ctx context.Context, msg mediator.Message) (any, error) {
	typedMessage := msg.(*GetEntriesByTypeMessage)

	return h.repo.GetByType(ctx, typedMessage.store, typedMessage.entryType)
}

func (h *GetEntriesByTypeHandler) GetType() string {
	return getEntriesByTypeType
}
