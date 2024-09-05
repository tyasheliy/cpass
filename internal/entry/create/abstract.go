package create

import (
	"context"
	"errors"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/passcl"
)

type Command interface {
	Create(ctx context.Context) error
}

type CommandFactory interface {
	Create(entry entry.Entry, data any) (Command, error)
}

type CommandFactoryImpl struct {
	client passcl.Client
}

func NewCommandFactoryImpl(client passcl.Client) *CommandFactoryImpl {
	return &CommandFactoryImpl{client: client}
}

func (f *CommandFactoryImpl) Create(e entry.Entry, data any) (Command, error) {
	switch typed := e.(type) {
	case *entry.PasswordEntry:
		strData, ok := data.(string)
		if !ok {
			return nil, errors.New("string data required")
		}

		return newPasswordCommand(f.client, typed, strData), nil
	case *entry.OtpEntry:
		uriData, ok := data.(string)
		if !ok {
			return nil, errors.New("string uri required")
		}

		return newOtpCommand(f.client, typed, uriData), nil
	case *entry.TodoEntry:
		linesData, ok := data.([]string)
		if !ok {
			return nil, errors.New("lines as string slice required")
		}

		return newTodoCommand(f.client, typed, linesData), nil
	default:
		return nil, errors.New("unknown entry type")
	}
}
