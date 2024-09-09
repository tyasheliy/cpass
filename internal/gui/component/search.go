package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Search struct {
	onInputChange func(text string) error
}

func NewSearch() *Search {
	return &Search{
		onInputChange: nil,
	}
}

func (c *Search) OnInputChange(handler func(text string) error) {
	c.onInputChange = handler
}

func (c *Search) Draw() (fyne.CanvasObject, error) {
	input := widget.NewEntry()
	input.SetPlaceHolder("Search for entry...")

	input.OnChanged = func(text string) {
		_ = c.onInputChange(text)
	}

	return input, nil
}
