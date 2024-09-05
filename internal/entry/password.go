package entry

const passwordSuffix = "pass"

type PasswordEntry struct {
	fileEntry
}

func NewPasswordEntry(parent *DirEntry, fileName string) *PasswordEntry {
	return &PasswordEntry{
		fileEntry{
			parent:   parent,
			fileName: fileName,
		},
	}
}
