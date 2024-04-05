package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()
	window := app.NewWindow("salam")
	window.SetContent(widget.NewLabel("tests"))

	window.ShowAndRun()
}
