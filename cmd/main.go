package main

import (
	"context"
	"fmt"
	"github.com/tyasheliy/cpass/internal/entry"
	"log"
	"os"
	"path/filepath"
)

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	storePath := filepath.Join(homeDir, ".password-store")

	m := entry.NewQueryManager(storePath)

	p := entry.NewDirEntry("test", nil)
	dir := entry.NewDirEntry("test1", p)

	entries, err := m.GetDirEntryChildren(context.Background(), dir)
	if err != nil {
		fmt.Println(err)
	}

	for _, e := range entries {
		fmt.Println()
	}
}
