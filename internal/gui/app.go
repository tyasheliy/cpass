package gui

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/tyasheliy/cpass/internal/gui/templates"
	"github.com/tyasheliy/cpass/internal/gui/templates/abstract"
)

type App struct {
	fyneApp           fyne.App
	window            fyne.Window
	templateAssembler abstract.TemplateAssembler
}

func NewApp(
	templateAssembler abstract.TemplateAssembler,
) *App {
	fyneApp := app.New()
	window := fyneApp.NewWindow("cpass")

	return &App{
		fyneApp:           fyneApp,
		window:            window,
		templateAssembler: templateAssembler,
	}
}

func (a *App) Run(ctx context.Context) error {
	mainTemplate := templates.NewMainTemplateVisitor()
	content, err := a.templateAssembler.BuildTemplate(ctx, mainTemplate)
	if err != nil {
		return err
	}

	a.window.SetContent(content)

	a.window.Show()
	a.fyneApp.Run()

	return nil
}
