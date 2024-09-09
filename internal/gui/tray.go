package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

type tray struct {
	a       *App
	mainWin fyne.Window
}

func newTray(a *App, mainWin fyne.Window) *tray {
	return &tray{
		a:       a,
		mainWin: mainWin,
	}
}

func (t *tray) init() {
	desk, ok := t.a.fyneApp.(desktop.App)
	if !ok {
		return
	}

	t.mainWin.SetCloseIntercept(func() {
		t.mainWin.Hide()
	})

	m := fyne.NewMenu("cpass",
		fyne.NewMenuItem("Show", func() {
			t.mainWin.Show()
		}),
		fyne.NewMenuItem("Quit", func() {
			t.a.fyneApp.Quit()
		}),
	)

	desk.SetSystemTrayMenu(m)
}
