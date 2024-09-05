package create

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
)

type todoCommand struct {
	client passcl.Client
	entry  *entry.TodoEntry
	lines  []string
}

func newTodoCommand(client passcl.Client, entry *entry.TodoEntry, lines []string) *todoCommand {
	return &todoCommand{client: client, entry: entry, lines: lines}
}

func (c *todoCommand) Create(ctx context.Context) error {
	passName := entry.PassName(c.entry)
	options := passcl.InsertOptions{
		Force:     false,
		MultiLine: true,
	}

	return c.client.Insert(
		ctx,
		passName,
		c.lines,
		options,
	)
}
