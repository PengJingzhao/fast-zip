package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

var (
	progressBar       *widget.ProgressBar
	selectedFileLabel *widget.Label
	targetFileLabel   *widget.Label
	restTimeFileLabel *widget.Label
	myApp fyne.App
)

func main() {
	myApp = app.New()
	CreateMainWindow()
}
