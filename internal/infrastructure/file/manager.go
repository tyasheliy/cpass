package file

import (
	"github.com/tyasheliy/cpass/internal/entity"
	"strings"
)

type Manager struct {
	rootPath string
}

func NewManager(rootPath string) *Manager {
	return &Manager{
		rootPath: rootPath,
	}
}

func (m *Manager) GetStoreFullName(store *entity.Store) string {
	names := make([]string, 0)
	names = append(names, store.Name)

	addStoreParentName(&names, store)

	namesLen := len(names)
	reversedNames := make([]string, namesLen)

	for i := range names {
		name := names[i]

		reversedNames[namesLen-1-i] = name
	}

	return strings.Join(reversedNames, "/")
}

func addStoreParentName(names *[]string, store *entity.Store) {
	if store.Parent == nil {
		return
	}

	*names = append(*names, store.Parent.Name)
	addStoreParentName(names, store.Parent)
}
