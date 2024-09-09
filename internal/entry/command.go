package entry

import "github.com/tyasheliy/cpass/internal/passcl"

type CommandManager interface {
}

type CommandManagerImpl struct {
	client passcl.Client
}
