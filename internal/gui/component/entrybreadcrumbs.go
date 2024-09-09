package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tyasheliy/cpass/internal/entry"
)

type EntryBreadcrumbs struct {
	onBreadcrumbClick func(e entry.Entry) error
	cont              *fyne.Container
}

func NewEntryBreadcrumbs(e entry.Entry) *EntryBreadcrumbs {
	cont := container.NewHBox()

	breadcrumbs := EntryBreadcrumbs{
		onBreadcrumbClick: nil,
		cont:              cont,
	}

	breadcrumbs.DrawEntry(e)

	return &breadcrumbs
}

func (c *EntryBreadcrumbs) OnBreadcrumbClick(handler func(e entry.Entry) error) {
	c.onBreadcrumbClick = handler
}

func (c *EntryBreadcrumbs) Draw() (fyne.CanvasObject, error) {
	return c.cont, nil
}

func (c *EntryBreadcrumbs) DrawEntry(e entry.Entry) {
	c.cont.RemoveAll()

	if e == nil {
		return
	}

	parents := entry.SplitEntryParents(e)

	for _, parent := range parents.Slice() {
		c.cont.Add(c.drawBreadcrumb(parent))
		c.cont.Add(widget.NewLabel("/"))
	}

	c.cont.Refresh()
}

func (c *EntryBreadcrumbs) drawBreadcrumb(e entry.Entry) fyne.CanvasObject {
	return widget.NewButton(e.Name(), func() {
		_ = c.onBreadcrumbClick(e)
	})
}
