package templates

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

type MainTemplateVisitor struct{}

func NewMainTemplateVisitor() *MainTemplateVisitor {
	return &MainTemplateVisitor{}
}

func (v *MainTemplateVisitor) GetCanvas(ctx context.Context, md mediator.MessageMediator) (fyne.CanvasObject, error) {
	return widget.NewLabel("main"), nil
}
