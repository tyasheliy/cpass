package entry

type DirEntry struct {
	name   string
	parent *DirEntry
}

func NewDirEntry(name string, parent *DirEntry) *DirEntry {
	return &DirEntry{
		name:   name,
		parent: parent,
	}
}

func (d *DirEntry) FileName() string {
	return d.name
}

func (d *DirEntry) Name() string {
	return d.name
}

func (d *DirEntry) Parent() *DirEntry {
	return d.parent
}
