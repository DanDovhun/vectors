package funcs

import "math"

type Angles struct {
	Radians float64
	Degrees float64
}

func ToDegrees(rad float64) float64 {
	return 180 * rad / math.Pi
}
