package extra

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Alert(appl fyne.App, home fyne.Window, error string) {
	alert := appl.NewWindow("Error")

	alert.SetContent(container.NewVBox(
		widget.NewLabel(error),
		widget.NewButton("OK", func() {
			alert.Hide()
		}),
	))

	alert.Resize(fyne.NewSize(175, 0))
	alert.Show()
}
