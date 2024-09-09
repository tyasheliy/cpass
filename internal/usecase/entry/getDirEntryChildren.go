package entry

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

const getDirEntryChildrenType = "getDirEntryChildren"

type GetDirEntryChildrenMessage struct {
	dirEntry *entry.DirEntry
}

func NewGetDirEntryChildrenMessage(dirEntry *entry.DirEntry) *GetDirEntryChildrenMessage {
	return &GetDirEntryChildrenMessage{dirEntry: dirEntry}
}

func (m *GetDirEntryChildrenMessage) GetHandlerType() string {
	return getDirEntryChildrenType
}

type GetDirEntryChildrenHandler struct {
	query *entry.QueryManager
}

func NewGetDirEntryChildrenHandler(query *entry.QueryManager) *GetDirEntryChildrenHandler {
	return &GetDirEntryChildrenHandler{query: query}
}

func (h *GetDirEntryChildrenHandler) Handle(ctx context.Context, msg mediator.Message) (any, error) {
	typedMsg := msg.(*GetDirEntryChildrenMessage)
	return h.query.GetDirEntryChildren(ctx, typedMsg.dirEntry)
}

func (h *GetDirEntryChildrenHandler) GetType() string {
	return getDirEntryChildrenType
}
