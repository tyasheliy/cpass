package component

import "fyne.io/fyne/v2"

type Component interface {
	Draw() (fyne.CanvasObject, error)
}
