package windows

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Home() {
	myApp := app.New()
	myWin := myApp.NewWindow("Vector Calculator")

	content := container.NewVBox(
		widget.NewButton("2D Vector", func() {
			TwoDim(myApp, myWin)
		}),

		widget.NewButton("3D Vector", func() {
			ThreeDim(myApp, myWin)
		}),
	)

	myWin.Resize(fyne.NewSize(275, 0))
	myWin.SetContent(content)
	myWin.ShowAndRun()
}
