package create

import (
	"context"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
)

type otpCommand struct {
	client passcl.Client
	entry  *entry.OtpEntry
	uri    string
}

func newOtpCommand(client passcl.Client, entry *entry.OtpEntry, uri string) *otpCommand {
	return &otpCommand{client: client, entry: entry, uri: uri}
}

func (c *otpCommand) Create(ctx context.Context) error {
	passName := entry.PassName(c.entry)
	options := passcl.InsertOtpOptions{
		Force: false,
	}

	return c.client.InsertOtp(
		ctx,
		passName,
		c.uri,
		options,
	)
}
