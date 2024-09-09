package template

import "fyne.io/fyne/v2"

type Template interface {
	Show(on fyne.Window) error
}
