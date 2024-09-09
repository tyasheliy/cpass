package component

import (
	"context"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/tyasheliy/cpass/internal/entry"
	entry2 "github.com/tyasheliy/cpass/internal/usecase/entry"
	"github.com/tyasheliy/cpass/internal/usecase/mediator"
	"os"
	"sync"
)

type EntryList struct {
	md              mediator.MessageMediator
	notif           *Notification
	cont            *fyne.Container
	w               fyne.Window
	onDirEntryClick func(dirEntry *entry.DirEntry) error
}

func NewEntryList(
	aggregate *entry.Aggregate,
	md mediator.MessageMediator,
	notif *Notification,
	w fyne.Window,
) (*EntryList, error) {
	vBox := container.NewVBox()

	entryList := EntryList{
		w:     w,
		md:    md,
		notif: notif,
		cont:  vBox,
	}

	err := entryList.DrawEntries(aggregate)
	if err != nil {
		return nil, err
	}

	return &entryList, nil
}

func (c *EntryList) OnDirEntryClick(handler func(dirEntry *entry.DirEntry) error) {
	c.onDirEntryClick = handler
}

func (c *EntryList) Draw() (fyne.CanvasObject, error) {
	return c.cont, nil
}

func (c *EntryList) DrawEntries(aggregate *entry.Aggregate) error {
	c.cont.RemoveAll()

	sortedEntries := aggregate.SortByPassName().Slice()
	dirEntriesIndexes := make([]int, 0)
	fileEntriesIndexes := make([]int, 0)

	for i := range sortedEntries {
		e := sortedEntries[i]

		switch e.(type) {
		case *entry.DirEntry:
			dirEntriesIndexes = append(dirEntriesIndexes, i)
		default:
			fileEntriesIndexes = append(fileEntriesIndexes, i)
		}
	}

	dirEntries := make([]*entry.DirEntry, len(dirEntriesIndexes))
	fileEntries := make([]entry.Entry, len(fileEntriesIndexes))

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i, entryI := range dirEntriesIndexes {
			dirEntries[i] = sortedEntries[entryI].(*entry.DirEntry)
		}
	}()

	go func() {
		defer wg.Done()
		for i, entryI := range fileEntriesIndexes {
			fileEntries[i] = sortedEntries[entryI]
		}
	}()

	wg.Wait()

	for _, dirEntry := range dirEntries {
		btn := widget.NewButton(fmt.Sprintf("%s/", dirEntry.Name()), func() {
			if c.onDirEntryClick != nil {
				err := c.onDirEntryClick(dirEntry)
				if err != nil {
					return
				}
			}

			rawChildren, err := c.md.Send(
				context.Background(),
				entry2.NewGetDirEntryChildrenMessage(dirEntry),
			)
			if err != nil {
				return
			}

			children := rawChildren.(*entry.Aggregate)
			err = c.DrawEntries(children)
		})

		c.cont.Add(btn)
	}

	ctx := context.Background()

	rawQuery, err := c.md.Send(ctx, entry2.NewGetEntryQueryManagerMessage())
	if err != nil {
		return err
	}
	query := rawQuery.(*entry.QueryManager)
	clipboard := c.w.Clipboard()

	for _, fileEntry := range fileEntries {
		var handler func()

		switch typed := fileEntry.(type) {
		case *entry.PasswordEntry:
			handler = func() {
				password, err := query.GetPassword(ctx, typed)
				if err != nil {
					dialog.ShowInformation("err", os.Getenv("PATH"), c.w)
					return
				}

				clipboard.SetContent(password)
				_ = c.notif.Show(fmt.Sprintf("Copied %s!", typed.Name()))
			}
		case *entry.OtpEntry:
			handler = func() {
				otp, err := query.GetOtp(ctx, typed)
				if err != nil {
					return
				}

				clipboard.SetContent(otp)
			}
		case *entry.TodoEntry:
			handler = func() {
				lines, err := query.GetTodoLines(ctx, typed)
				if err != nil {
					return
				}

				fmt.Println(lines)
			}
		}

		btn := widget.NewButton(fileEntry.Name(), handler)

		c.cont.Add(btn)
	}

	c.cont.Refresh()
	return nil
}
