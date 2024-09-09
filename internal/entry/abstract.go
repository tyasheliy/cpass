package entry

import (
	"path/filepath"
	"strings"
)

type Entry interface {
	FileName() string
	Name() string
	Parent() *DirEntry
}

type fileEntry struct {
	parent   *DirEntry
	fileName string
}

func (p *fileEntry) FileName() string {
	return p.fileName
}

func (p *fileEntry) Name() string {
	split := strings.Split(p.fileName, ".")

	return strings.Join(split[:len(split)-1], ".")
}

func (p *fileEntry) Parent() *DirEntry {
	return p.parent
}

func Depth(entry Entry) int {
	if entry.Parent() == nil {
		return 0
	}

	return 1 + Depth(entry.Parent())
}

func PassName(entry Entry) string {
	if entry.Parent() == nil {
		return entry.Name()
	}

	return filepath.Join(PassName(entry.Parent()), entry.Name())
}

func FullFileName(entry Entry) string {
	if entry.Parent() == nil {
		return entry.FileName()
	}

	return filepath.Join(FullFileName(entry.Parent()), entry.FileName())
}

func SplitEntryParents(entry Entry) *Aggregate {
	return NewAggregate(splitEntryParents(entry)...)
}

func splitEntryParents(entry Entry) []Entry {
	if entry.Parent() == nil {
		return []Entry{entry}
	}

	return append(splitEntryParents(entry.Parent()), entry)
}
