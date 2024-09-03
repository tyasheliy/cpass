package abstract

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/tyasheliy/cpass/internal/gui/components"
	"github.com/tyasheliy/cpass/internal/gui/visitor"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
)

type TemplateAssembler interface {
	BuildTemplate(ctx context.Context, visitor visitor.CanvasVisitor) (fyne.CanvasObject, error)
}

type TemplateAssemblerImpl struct {
	messageMediator mediator.MessageMediator
}

func NewTemplateAssembler(messageMediator mediator.MessageMediator) *TemplateAssemblerImpl {
	return &TemplateAssemblerImpl{
		messageMediator: messageMediator,
	}
}

func (a *TemplateAssemblerImpl) BuildTemplate(ctx context.Context, visitor visitor.CanvasVisitor) (fyne.CanvasObject, error) {
	header, err := components.NewHeaderVisitor().GetCanvas(ctx, a.messageMediator)
	if err != nil {
		return nil, err
	}

	content, err := visitor.GetCanvas(ctx, a.messageMediator)
	if err != nil {
		return nil, err
	}

	return container.NewVBox(header, content), nil
}
