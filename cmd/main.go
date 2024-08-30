package main

import (
	"context"
	"fmt"
	"github.com/tyasheliy/cpass/internal/entity"
	"github.com/tyasheliy/cpass/internal/infrastructure/file"
	"github.com/tyasheliy/cpass/internal/infrastructure/repository/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
)

func main() {
	m := file.NewManager("/Users/tyasheliy/.password-store")
	c := passcl.NewOsClient()
	r := entry.NewFileEntryRepository(m, c)

	store := &entity.Store{
		Name:   "test",
		Parent: nil,
	}

	fmt.Println(r.Delete(context.Background(), store, "first"))
}
