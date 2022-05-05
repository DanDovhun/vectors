package windows

import (
	"Vectors/dimensions/two"
	"Vectors/extra"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func TwoDim(appl fyne.App, home fyne.Window) {
	win := appl.NewWindow("2D Vectors")
	win.SetMaster()
	home.Hide()

	aXInp := widget.NewEntry()
	aYInp := widget.NewEntry()

	aXInp.SetPlaceHolder("X")
	aYInp.SetPlaceHolder("B")

	bXInp := widget.NewEntry()
	bYInp := widget.NewEntry()

	bXInp.SetPlaceHolder("X")
	bYInp.SetPlaceHolder("B")

	content := container.NewVBox(
		widget.NewLabel("Vector A"),
		aXInp,
		aYInp,

		widget.NewLabel("Vector B"),
		bXInp,
		bYInp,

		widget.NewButton("Calculate", func() {
			if len(aXInp.Text) == 0 || len(aYInp.Text) == 0 || len(bXInp.Text) == 0 || len(bYInp.Text) == 0 {
				extra.Alert(appl, win, "Please enter something")
				return
			}

			xA, err1 := strconv.ParseFloat(aXInp.Text, 64)
			yA, err2 := strconv.ParseFloat(aYInp.Text, 64)

			xB, err3 := strconv.ParseFloat(bXInp.Text, 64)
			yB, err4 := strconv.ParseFloat(bYInp.Text, 64)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				extra.Alert(appl, win, "Please only enter numbers")
				return
			}

			vecA := two.Vector{X: xA, Y: yA}
			vecB := two.Vector{X: xB, Y: yB}

			extra.TwoDim(appl, win, vecA, vecB)
		}),

		widget.NewButton("Back Home", func() {
			win.Hide()
			home.Show()
			home.SetMaster()
		}),
	)

	win.Resize(fyne.NewSize(275, 0))
	win.SetContent(content)
	win.Show()
}
