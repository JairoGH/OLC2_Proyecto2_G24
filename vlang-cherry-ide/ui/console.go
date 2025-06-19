package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Console struct {
	Container fyne.CanvasObject
	Output    *widget.Entry
}

func NewConsole() *Console {

	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("La salida del programa aparecerá aquí...")

	output.TextStyle = fyne.TextStyle{Monospace: true}

	label := widget.NewLabel("Consola")

	console := &Console{
		Output: output,
	}

	scrollOutput := container.NewScroll(output)
	console.Container = container.NewBorder(
		label,        // top
		nil,          // bottom
		nil,          // left
		nil,          // right
		scrollOutput, // center
	)

	return console
}

func (c *Console) AddOutput(text string) {
	current := c.Output.Text
	c.Output.SetText(current + text)

	c.Output.CursorRow = len(c.Output.Text)
}

func (c *Console) ClearOutput() {
	c.Output.SetText("")
}

func (c *Console) SetConsoleStyle() {

	c.Output.TextStyle = fyne.TextStyle{
		Monospace: true,
		Bold:      false,
	}
	c.Output.Refresh()
}
