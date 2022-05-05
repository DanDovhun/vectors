package windows

import (
	"Vectors/dimensions/three"
	"Vectors/extra"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func ThreeDim(appl fyne.App, home fyne.Window) {
	win := appl.NewWindow("2D Vectors")
	win.SetMaster()
	home.Hide()

	aXInp := widget.NewEntry()
	aYInp := widget.NewEntry()
	aZInp := widget.NewEntry()

	aXInp.SetPlaceHolder("X")
	aYInp.SetPlaceHolder("B")
	aZInp.SetPlaceHolder("Z")

	bXInp := widget.NewEntry()
	bYInp := widget.NewEntry()
	bZInp := widget.NewEntry()

	bXInp.SetPlaceHolder("X")
	bYInp.SetPlaceHolder("B")
	bZInp.SetPlaceHolder("Z")

	content := container.NewVBox(
		widget.NewLabel("Vector A"),
		aXInp,
		aYInp,
		aZInp,

		widget.NewLabel("Vector B"),
		bXInp,
		bYInp,
		bZInp,

		widget.NewButton("Calculate", func() {
			if len(aXInp.Text) == 0 || len(aYInp.Text) == 0 || len(aZInp.Text) == 0 || len(bXInp.Text) == 0 || len(bYInp.Text) == 0 || len(bZInp.Text) == 0 {
				extra.Alert(appl, win, "Please enter something")
				return
			}

			xA, err1 := strconv.ParseFloat(aXInp.Text, 64)
			yA, err2 := strconv.ParseFloat(aYInp.Text, 64)
			zA, err3 := strconv.ParseFloat(aZInp.Text, 64)

			xB, err4 := strconv.ParseFloat(bXInp.Text, 64)
			yB, err5 := strconv.ParseFloat(bYInp.Text, 64)
			zB, err6 := strconv.ParseFloat(bZInp.Text, 64)

			if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil {
				extra.Alert(appl, win, "Please only enter numbers")
				return
			}

			vecA := three.Vector{X: xA, Y: yA, Z: zA}
			vecB := three.Vector{X: xB, Y: yB, Z: zB}

			extra.ThreeDim(appl, vecA, vecB)
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
