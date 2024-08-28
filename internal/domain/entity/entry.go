package entity

type Entry struct {
	store     *Store
	name      string
	entryType EntryType
}

type EntryType int

const (
	PasswordEntryType EntryType = iota
	OtpEntryType      EntryType = iota
	TodoEntryType     EntryType = iota
)

type EntryRepository interface {
	Create(entry *Entry) error
	Get() ([]*Entry, error)
	GetByType(store *Store, t EntryType) ([]*Entry, error)
	GetByName(store *Store, name string) (*Entry, error)
	Delete(store *Store, name string) error
}
