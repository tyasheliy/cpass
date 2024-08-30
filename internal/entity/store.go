package entity

import "context"

type Store struct {
	Name   string
	Parent *Store
}

type StoreRepository interface {
	Create(ctx context.Context, store *Store) error
	Delete(ctx context.Context, store *Store) error
}
