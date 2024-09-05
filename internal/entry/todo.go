package entry

const todoSuffix = "todo"

type TodoEntry struct {
	fileEntry
}

func NewTodoEntry(parent *DirEntry, fileName string) *TodoEntry {
	return &TodoEntry{
		fileEntry{
			parent:   parent,
			fileName: fileName,
		},
	}
}
