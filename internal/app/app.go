package app

import (
	"github.com/tyasheliy/cpass/internal/logger"
)

type App struct {
	logger logger.Logger
}

func NewApp(
	logger logger.Logger,
) *App {
	return &App{
		logger,
		renderer,
	}
}
