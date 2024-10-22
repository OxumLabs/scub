package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	scub := app.New()
	println(scub)

	// Create a new window
	swin := scub.NewWindow("SCUB - Engine")
	swin.Resize(fyne.NewSize(720, 480))
	swin.SetContent(container.NewVBox(
		widget.NewLabel("Hello, World!"),
		widget.NewButton("Quit", func() {
			scub.Quit()
		}),
	))

	// Show and run the application
	swin.ShowAndRun()
}
