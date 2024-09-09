package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/tyasheliy/cpass/internal/gui/template"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

type App struct {
	md      mediator.MessageMediator
	fyneApp fyne.App
}

func NewApp(md mediator.MessageMediator) *App {
	a := app.NewWithID("com.tyasheliy.cpass")

	return &App{
		fyneApp: a,
		md:      md,
	}
}

func (a *App) Run() error {
	w := a.fyneApp.NewWindow("cpass")

	mainTemplate := template.NewMain(a.md)
	err := mainTemplate.Show(w)
	if err != nil {
		return err
	}

	appTray := newTray(a, w)
	appTray.init()

	a.fyneApp.Run()
	return nil
}
