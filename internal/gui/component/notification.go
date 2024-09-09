package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

type Notification struct {
	cont        fyne.CanvasObject
	textBinding binding.String
	hideTimer   *time.Timer
	dur         time.Duration
}

func NewNotification() *Notification {
	textBinding := binding.NewString()

	notif := Notification{
		textBinding: textBinding,
		hideTimer:   nil,
		dur:         5 * time.Second,
	}

	notif.drawCont()
	notif.cont.Hide()

	return &notif
}

func (n *Notification) drawCont() {
	label := widget.NewLabelWithData(n.textBinding)
	label.TextStyle.Bold = true
	cont := container.NewWithoutLayout(label)

	n.cont = cont
}

func (n *Notification) Show(text string) error {
	err := n.textBinding.Set(text)
	if err != nil {
		return err
	}

	go func() {
		if n.hideTimer != nil {
			n.hideTimer.Reset(n.dur)
		} else {
			n.hideTimer = time.AfterFunc(n.dur, func() {
				n.cont.Hide()
			})
		}
	}()

	n.cont.Show()
	return nil
}

func (n *Notification) Draw() (fyne.CanvasObject, error) {
	return n.cont, nil
}
