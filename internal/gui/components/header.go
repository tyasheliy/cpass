package components

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

type HeaderVisitor struct{}

func NewHeaderVisitor() *HeaderVisitor {
	return &HeaderVisitor{}
}

func (v *HeaderVisitor) GetCanvas(ctx context.Context, mediator mediator.MessageMediator) (fyne.CanvasObject, error) {
	leftLabel := widget.NewLabel("left")
	rightLabel := widget.NewLabel("right")

	content := container.NewHBox(leftLabel, layout.NewSpacer(), rightLabel)

	return content, nil
}
