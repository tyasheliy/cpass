package entry

import (
	"context"
	"github.com/tyasheliy/cpass/internal/passcl"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type QueryManager struct {
	client   passcl.Client
	rootPath string
}

func NewQueryManager(rootPath string, client passcl.Client) *QueryManager {
	return &QueryManager{
		client:   client,
		rootPath: rootPath,
	}
}

func (m *QueryManager) GetPassword(ctx context.Context, entry *PasswordEntry) (string, error) {
	passName := PassName(entry)
	return m.client.Show(ctx, passName)
}

func (m *QueryManager) GetOtp(ctx context.Context, entry *OtpEntry) (string, error) {
	passName := PassName(entry)
	return m.client.ShowOtp(ctx, passName)
}

func (m *QueryManager) GetTodoLines(ctx context.Context, entry *TodoEntry) ([]string, error) {
	passName := PassName(entry)

	rawLines, err := m.client.Show(ctx, passName)
	if err != nil {
		return nil, err
	}

	return strings.Split(rawLines, "\n"), nil
}

func (m *QueryManager) GetDirEntryChildren(ctx context.Context, entry *DirEntry) (*Aggregate, error) {
	root := m.getEntryPathFromRoot(entry)

	entries := make([]Entry, 0)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		rel, err := m.getRelPathFromDirEntry(entry, path)
		if err != nil {
			return err
		}

		if strings.HasPrefix(rel, ".") {
			return nil
		}

		entries = append(entries, m.getEntryByPath(entry, rel))

		return nil
	})
	if err != nil {
		return nil, err
	}

	return NewAggregate(entries...), nil
}

func (m *QueryManager) getEntryByPath(rel *DirEntry, path string) Entry {
	splitPath := strings.Split(path, GetOsSeparatorAsStr())
	pathLastIndex := len(splitPath) - 1

	var parent *DirEntry
	if rel != nil {
		parent = rel
	} else {
		parent = nil
	}

	for _, name := range splitPath[:pathLastIndex] {
		parent = &DirEntry{
			name:   name,
			parent: parent,
		}
	}

	fileName := splitPath[pathLastIndex]
	splitName := strings.Split(fileName, ".")
	nameLastIndex := len(splitName) - 1

	if splitName[nameLastIndex] != "gpg" {
		return &DirEntry{
			name:   fileName,
			parent: parent,
		}
	}

	switch splitName[nameLastIndex-1] {
	case passwordSuffix:
		return NewPasswordEntry(parent, fileName)
	case otpSuffix:
		return NewOtpEntry(parent, fileName)
	case todoSuffix:
		return NewTodoEntry(parent, fileName)
	default:
		return NewPasswordEntry(parent, fileName)
	}
}

func (m *QueryManager) getRelPathFromDirEntry(entry *DirEntry, path string) (string, error) {
	return filepath.Rel(m.getEntryPathFromRoot(entry), path)
}

func (m *QueryManager) getEntryPathFromRoot(entry Entry) string {
	return m.getFromRoot(FullFileName(entry))
}

func (m *QueryManager) getFromRoot(path string) string {
	return filepath.Join(m.rootPath, path)
}

func GetOsSeparatorAsStr() string {
	return string(os.PathSeparator)
}
