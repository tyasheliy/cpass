package view

type View interface {
	Render() error
}

type ViewFactory interface {
	CreateWindow(title string, content View) Window
	CreateLabel(text string) View
	CreateRow(children []View) View
	CreateCol(children []View) View
}

type Window interface {
	Render() error
	SetTitle(title string)
	SetContent(content View)
}

type Label interface {
	Render() error
	SetText(text string)
}

type Row interface {
}
