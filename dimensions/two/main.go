package two

import (
	"Vectors/funcs"
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
}

func (vec Vector) ToString() string {
	return fmt.Sprintf("(%v, %v)", vec.X, vec.Y)
}

// Returns float64

func (vec Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2))
}

func (vec Vector) DotProduct(vecB Vector) float64 {
	return vec.X*vecB.X + vec.Y*vecB.Y
}

func (vec Vector) CrossProduct(vecB Vector) float64 {
	return vec.X*vecB.Y - vec.Y*vecB.X
}

// Returns Vector

func (vec Vector) Add(vecB Vector) Vector {
	return Vector{
		X: vec.X + vecB.X,
		Y: vec.Y + vecB.Y,
	}
}
func (vec Vector) Sub(vecB Vector) Vector {
	return Vector{
		X: vec.X - vecB.X,
		Y: vec.Y - vecB.Y,
	}
}

// Returns Angle

func (vec Vector) Elevation() funcs.Angles {
	rad := math.Atan(vec.Y / vec.X)
	deg := funcs.ToDegrees(rad)

	return funcs.Angles{
		Radians: funcs.Round(rad, 5),
		Degrees: funcs.Round(deg, 5),
	}
}

func (vec Vector) AngleBetween(vecB Vector) funcs.Angles {
	rad := math.Acos(vec.CrossProduct(vecB) / (vec.Magnitude() * vecB.Magnitude()))
	deg := funcs.ToDegrees(rad)

	return funcs.Angles{
		Radians: funcs.Round(rad, 5),
		Degrees: funcs.Round(deg, 5),
	}
}
