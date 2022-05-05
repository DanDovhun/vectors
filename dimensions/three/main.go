package three

import (
	"Vectors/dimensions/two"
	"Vectors/funcs"
	"fmt"
	"math"
)

type Vector struct {
	X float64
	Y float64
	Z float64
}

func (vec Vector) ToString() string {
	return fmt.Sprintf("(%v, %v, %v)", vec.X, vec.Y, vec.Z)
}

// Returns float64

func (vec Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(vec.X, 2) + math.Pow(vec.Y, 2) + math.Pow(vec.Z, 2))
}

func (vec Vector) DotProduct(vecB Vector) float64 {
	return vec.X*vecB.X + vec.Y*vecB.Y + vec.Z*vecB.Z
}

// Returns Vector

func (vec Vector) Add(vecB Vector) Vector {
	return Vector{
		X: vec.X + vecB.X,
		Y: vec.Y + vecB.Y,
		Z: vec.Z + vecB.Z,
	}
}

func (vec Vector) Sub(vecB Vector) Vector {
	return Vector{
		X: vec.X - vecB.X,
		Y: vec.Y - vecB.Y,
		Z: vec.Z - vecB.Z,
	}
}

func (vec Vector) CrossProduct(vecB Vector) Vector {
	return Vector{
		X: vec.Y*vecB.Z - vec.Z*vecB.Y,
		Y: vec.Z*vecB.X - vec.X*vecB.Z,
		Z: vec.X*vecB.Y - vec.Y*vecB.X,
	}
}

// Returns Angle

func (vec Vector) Elevation() funcs.Angles {
	twoDim := two.Vector{X: vec.X, Y: vec.Z}.Magnitude()
	radians := math.Atan(vec.Y / twoDim)
	degrees := funcs.ToDegrees(radians)

	return funcs.Angles{
		Radians: funcs.Round(radians, 5),
		Degrees: funcs.Round(degrees, 5),
	}
}

func (vec Vector) Rotation() funcs.Angles {
	radians := math.Atan(vec.Z / vec.X)
	degrees := funcs.ToDegrees(radians)

	return funcs.Angles{
		Radians: funcs.Round(radians, 5),
		Degrees: funcs.Round(degrees, 5),
	}
}

func (vec Vector) AngleBetween(vecB Vector) funcs.Angles {
	radians := vec.DotProduct(vecB) / (vec.Magnitude() * vecB.Magnitude())
	degrees := funcs.ToDegrees(radians)

	return funcs.Angles{
		Radians: funcs.Round(radians, 5),
		Degrees: funcs.Round(degrees, 5),
	}
}
