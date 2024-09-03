package visitor

import (
	"context"
	"fyne.io/fyne/v2"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

type CanvasVisitor interface {
	GetCanvas(ctx context.Context, md mediator.MessageMediator) (fyne.CanvasObject, error)
}
