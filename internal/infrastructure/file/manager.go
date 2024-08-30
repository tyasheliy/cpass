package file

import (
	"github.com/tyasheliy/cpass/internal/entity"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	OTP_DIR  = "otp"
	TODO_DIR = "todo"
)

type Manager struct {
	RootPath string
}

func NewManager(rootPath string) *Manager {
	return &Manager{
		RootPath: rootPath,
	}
}

func (m *Manager) GetStorePath(store *entity.Store) string {
	names := make([]string, 0)
	names = append(names, store.Name)

	addStoreParentName(&names, store)

	namesLen := len(names)
	reversedNames := make([]string, namesLen)

	for i := range names {
		name := names[i]

		reversedNames[namesLen-1-i] = name
	}

	return filepath.Join(reversedNames...)
}

func addStoreParentName(names *[]string, store *entity.Store) {
	if store.Parent == nil {
		return
	}

	*names = append(*names, store.Parent.Name)
	addStoreParentName(names, store.Parent)
}

func (m *Manager) GetEntries(root string, typeFilter *entity.EntryType) ([]*entity.Entry, error) {
	entries := make([]*entity.Entry, 0)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		var rel string
		rel, err = filepath.Rel(m.RootPath, path)
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasPrefix(rel, ".") {
			return nil
		}

		entry, err := m.GetEntryByPath(path)
		if err != nil {
			return err
		}

		if typeFilter != nil && entry.EntryType == *typeFilter || typeFilter == nil {
			entries = append(entries, entry)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (m *Manager) GetEntryByPath(path string) (*entity.Entry, error) {
	rel, err := filepath.Rel(m.RootPath, path)
	if err != nil {
		return nil, err
	}

	dirPath, fileName := filepath.Split(rel)
	if len(dirPath) > 1 {
		dirPath = dirPath[:len(dirPath)-1]
	}
	splitDirPath := strings.Split(dirPath, string(os.PathSeparator))
	typeDir := splitDirPath[len(splitDirPath)-1]

	var entryType entity.EntryType
	var storeDirs *[]string

	switch typeDir {
	case OTP_DIR:
		storeDirs = m.withoutTypeDir(splitDirPath)
		entryType = entity.OtpEntryType
	case TODO_DIR:
		storeDirs = m.withoutTypeDir(splitDirPath)
		entryType = entity.TodoEntryType
	default:
		storeDirs = &splitDirPath
		entryType = entity.PasswordEntryType
	}

	var store *entity.Store
	store = nil

	for _, name := range *storeDirs {
		if name == "" {
			continue
		}

		store = &entity.Store{
			Name:   name,
			Parent: store,
		}
	}

	entryName := strings.Split(fileName, ".")[0]

	return &entity.Entry{
		Store:     store,
		Name:      entryName,
		EntryType: entryType,
	}, nil
}

func (m *Manager) withoutTypeDir(splitDirPath []string) *[]string {
	without := splitDirPath[:len(splitDirPath)-1]
	return &without
}
