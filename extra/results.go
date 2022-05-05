package extra

import (
	"Vectors/dimensions/three"
	"Vectors/dimensions/two"
	"Vectors/funcs"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ThreeDim(appl fyne.App, vecA, vecB three.Vector) {
	// Gets results
	// Vector A parameters
	vecAMag := vecA.Magnitude()
	vecAElv := vecA.Elevation()
	vecARot := vecA.Rotation()
	vecAForm := vecA.ToString()

	// Vector B parameters
	vecBMag := vecB.Magnitude()
	vecBElv := vecB.Elevation()
	vecBRot := vecB.Rotation()
	vecBForm := vecB.ToString()

	// Addition, subtraction, products and angle between
	added := vecA.Add(vecB)
	subtracted := vecA.Sub(vecB)
	dotProduct := vecA.DotProduct(vecB)
	crossProduct := vecA.CrossProduct(vecB)
	angleBetween := vecA.AngleBetween(vecB)

	wind := appl.NewWindow("Results")

	content := container.NewVBox(
		widget.NewLabel("Vector A"),
		widget.NewLabel(fmt.Sprintf("Formula: %v", vecAForm)),
		widget.NewLabel(fmt.Sprintf("Magnitude: %v", funcs.Round(vecAMag, 5))),
		widget.NewLabel(fmt.Sprintf("Elevation: %v radians; %v degrees", vecAElv.Radians, vecAElv.Degrees)),
		widget.NewLabel(fmt.Sprintf("Rotation: %v radians; %v degrees\n", vecARot.Radians, vecARot.Degrees)),

		widget.NewLabel("Vector B"),
		widget.NewLabel(fmt.Sprintf("Formula: %v", vecBForm)),
		widget.NewLabel(fmt.Sprintf("Magnitude: %v", funcs.Round(vecBMag, 5))),
		widget.NewLabel(fmt.Sprintf("Elevation: %v radians; %v degrees", vecBElv.Radians, vecBElv.Degrees)),
		widget.NewLabel(fmt.Sprintf("Rotation: %v radians; %v degrees\n", vecBRot.Radians, vecBRot.Degrees)),

		widget.NewLabel("Arithmetics"),
		widget.NewLabel(fmt.Sprintf("A + B = %v", added)),
		widget.NewLabel(fmt.Sprintf("A - B = %v", subtracted)),
		widget.NewLabel(fmt.Sprintf("Dot Product: %v", dotProduct)),
		widget.NewLabel(fmt.Sprintf("Cross Product: %v", crossProduct)),
		widget.NewLabel(fmt.Sprintf("Degrees Between: %v radians; %v degrees", angleBetween.Radians, angleBetween.Degrees)),
	)

	wind.SetContent(content)
	wind.Show()
}

func TwoDim(appl fyne.App, home fyne.Window, vecA, vecB two.Vector) {
	// Gets results
	// Vector A parameters
	vecAMag := vecA.Magnitude()
	vecAElv := vecA.Elevation()
	vecAForm := vecA.ToString()

	// Vector B parameters
	vecBMag := vecB.Magnitude()
	vecBElv := vecB.Elevation()
	vecBForm := vecB.ToString()

	// Addition, subtraction, products and angle between
	added := vecA.Add(vecB)
	subtracted := vecA.Sub(vecB)
	dotProduct := vecA.DotProduct(vecB)
	crossProduct := vecA.CrossProduct(vecB)
	angleBetween := vecA.AngleBetween(vecB)

	// Show Results
	wind := appl.NewWindow("Results")

	content := container.NewVBox(
		widget.NewLabel("Vector A"),
		widget.NewLabel(fmt.Sprintf("Formula: %v", vecAForm)),
		widget.NewLabel(fmt.Sprintf("Magnitude: %v", funcs.Round(vecAMag, 5))),
		widget.NewLabel(fmt.Sprintf("Elevation: %v radians; %v degrees\n", vecAElv.Radians, vecAElv.Degrees)),

		widget.NewLabel("Vector B"),
		widget.NewLabel(fmt.Sprintf("Formula: %v", vecBForm)),
		widget.NewLabel(fmt.Sprintf("Magnitude: %v", funcs.Round(vecBMag, 5))),
		widget.NewLabel(fmt.Sprintf("Elevation: %v radians; %v degrees\n", vecBElv.Radians, vecBElv.Degrees)),

		widget.NewLabel("Arithmetics"),
		widget.NewLabel(fmt.Sprintf("A + B = %v", added)),
		widget.NewLabel(fmt.Sprintf("A - B = %v", subtracted)),
		widget.NewLabel(fmt.Sprintf("Dot Product: %v", dotProduct)),
		widget.NewLabel(fmt.Sprintf("Cross Product: %v", crossProduct)),
		widget.NewLabel(fmt.Sprintf("Degrees Between: %v radians; %v degrees", angleBetween.Radians, angleBetween.Degrees)),
	)

	wind.SetContent(content)
	wind.Show()
}
