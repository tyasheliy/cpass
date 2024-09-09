package template

import (
	"context"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/tyasheliy/cpass/internal/entry"
	"github.com/tyasheliy/cpass/internal/gui/component"
	entry2 "github.com/tyasheliy/cpass/internal/usecase/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
	"strings"
)

type Main struct {
	md mediator.MessageMediator
}

func NewMain(md mediator.MessageMediator) *Main {
	return &Main{md: md}
}

func (m *Main) Show(on fyne.Window) error {
	on.Resize(fyne.NewSize(300, 400))

	entryListOverlay, err := m.setupEntryListOverlay(on)
	if err != nil {
		return err
	}

	on.SetContent(entryListOverlay)
	on.Show()
	return nil
}

func (m *Main) setupEntryListOverlay(on fyne.Window) (fyne.CanvasObject, error) {
	rootDirEntry := entry.NewDirEntry("", nil)
	rawChildren, err := m.md.Send(
		context.Background(),
		entry2.NewGetDirEntryChildrenMessage(rootDirEntry),
	)
	if err != nil {
		return nil, err
	}

	children := rawChildren.(*entry.Aggregate)

	notif := component.NewNotification()
	notifCanvas, _ := notif.Draw()

	entryList, err := component.NewEntryList(children, m.md, notif, on)
	if err != nil {
		return nil, err
	}

	entryListCanvas, err := entryList.Draw()
	if err != nil {
		return nil, err
	}

	entryListScroll := container.NewVScroll(entryListCanvas)

	breadcrumbs := component.NewEntryBreadcrumbs(nil)
	breadcrumbsCanvas, err := breadcrumbs.Draw()
	if err != nil {
		return nil, err
	}

	breadcrumbs.OnBreadcrumbClick(func(e entry.Entry) error {
		dirEntry, ok := e.(*entry.DirEntry)
		if !ok {
			return nil
		}

		rawChildren, err := m.md.Send(
			context.Background(),
			entry2.NewGetDirEntryChildrenMessage(dirEntry),
		)
		if err != nil {
			return err
		}

		children := rawChildren.(*entry.Aggregate)
		breadcrumbs.DrawEntry(dirEntry)
		return entryList.DrawEntries(children)
	})
	entryList.OnDirEntryClick(func(dirEntry *entry.DirEntry) error {
		breadcrumbs.DrawEntry(dirEntry)
		return nil
	})

	fullFileNamesMap := make(map[string]entry.Entry)
	for _, child := range children.Slice() {
		fileName := entry.FullFileName(child)
		fullFileNamesMap[fileName] = child
	}

	search := component.NewSearch()
	search.OnInputChange(func(text string) error {
		results := make([]entry.Entry, 0)

		for fileName, child := range fullFileNamesMap {
			if strings.HasPrefix(fileName, text) {
				results = append(results, child)
			}
		}

		aggregate := entry.NewAggregate(results...)
		go breadcrumbs.DrawEntry(nil)
		return entryList.DrawEntries(aggregate)
	})

	searchCanvas, err := search.Draw()
	if err != nil {
		return nil, err
	}

	topCont := container.NewVBox(searchCanvas, breadcrumbsCanvas)

	cont := container.NewBorder(topCont, notifCanvas, nil, nil, entryListScroll)
	return cont, nil
}
