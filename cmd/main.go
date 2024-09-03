package main

import (
	"context"
	"fmt"
	"github.com/tyasheliy/cpass/internal/entity"
	"github.com/tyasheliy/cpass/internal/infrastructure/file"
	"github.com/tyasheliy/cpass/internal/infrastructure/repository/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
	entry_handler "github.com/tyasheliy/cpass/internal/usecase/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
	"os"
	"path/filepath"
)

func main() {
	homeDir, _ := os.UserHomeDir()
	rootPath := filepath.FromSlash(fmt.Sprintf("%s/%s", homeDir, ".password-store"))

	fileManager := file.NewManager(rootPath)
	passClient := passcl.NewOsClient()

	entryRepo := entry.NewFileEntryRepository(fileManager, passClient)

	messageMediator := mediator.NewMessageMediator()

	h := entry_handler.NewGetEntriesByTypeHandler(entryRepo)
	_ = messageMediator.Register(h)

	store := entity.Store{
		Name:   "",
		Parent: nil,
	}

	msg := entry_handler.NewGetEntriesByTypeMessage(&store, entity.PasswordEntryType)

	rawEntries, err := messageMediator.Send(context.Background(), msg)
	if err != nil {
		panic(err)
	}

	entries := rawEntries.([]*entity.Entry)

	fmt.Println(entries[5])
}
