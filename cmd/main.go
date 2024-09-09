package main

import (
	"fmt"
	"github.com/haroldadmin/pathfix"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/gui"
	"github.com/tyasheliy/cpass/internal/passcl"
	entry2 "github.com/tyasheliy/cpass/internal/usecase/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
	"os"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	err = pathfix.Fix()
	if err != nil {
		panic(err)
	}

	storePath := filepath.Join(homeDir, ".password-store")

	cl, err := passcl.NewOsClient()
	if err != nil {
		panic(err)
	}

	query := entry.NewQueryManager(storePath, cl)

	messageMediator := mediator.NewMessageMediator()

	getQueryHandler := entry2.NewGetEntryQueryManagerHandler(query)
	messageMediator.Register(getQueryHandler)

	getDirEntryChildrenHandler := entry2.NewGetDirEntryChildrenHandler(query)
	messageMediator.Register(getDirEntryChildrenHandler)

	app := gui.NewApp(messageMediator)
	fmt.Println(app.Run())
}
