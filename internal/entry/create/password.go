package create

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
)

type passwordCommand struct {
	client   passcl.Client
	entry    *entry.PasswordEntry
	password string
}

func newPasswordCommand(client passcl.Client, entry *entry.PasswordEntry, password string) *passwordCommand {
	return &passwordCommand{
		client:   client,
		entry:    entry,
		password: password,
	}
}

func (c *passwordCommand) Create(ctx context.Context) error {
	passName := entry.PassName(c.entry)
	options := passcl.InsertOptions{
		Force:     false,
		MultiLine: false,
	}

	return c.client.Insert(
		ctx,
		passName,
		[]string{c.password},
		options,
	)
}
