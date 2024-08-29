package entity

import "context"

type Entry struct {
	Store     *Store
	Name      string
	EntryType EntryType
}

type EntryType int

const (
	PasswordEntryType EntryType = iota
	OtpEntryType      EntryType = iota
	TodoEntryType     EntryType = iota
)

type EntryRepository interface {
	CreatePassword(ctx context.Context, store *Store, name string, password string) (*Entry, error)
	CreateOtp(ctx context.Context, store *Store, name string, uri string) (*Entry, error)
	CreateTodo(ctx context.Context, store *Store, name string, lines []string) (*Entry, error)
	GeneratePassword(ctx context.Context, store *Store, name string, gen PasswordGeneration) (*Entry, error)
	Get(ctx context.Context) ([]*Entry, error)
	GetByStore(ctx context.Context, store *Store) ([]*Entry, error)
	GetByType(ctx context.Context, store *Store, t EntryType) ([]*Entry, error)
	GetByName(ctx context.Context, store *Store, name string) (*Entry, error)
	Delete(ctx context.Context, store *Store, name string) error
}
