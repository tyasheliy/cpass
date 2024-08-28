package entity

type Store struct {
	name   string
	key    string
	parent *Store
}

type StoreRepository interface {
	Create(store *Store) error
	Delete(parent *Store, name string) error
}
