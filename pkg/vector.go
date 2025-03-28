package pkg

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vector) ToColor() Color {
	return NewColor(v.X, v.Y, v.Z)
}

func (v Vector) AddVector(u Vector) Vector {
	return Vector{v.X + u.X, v.Y + u.Y, v.Z + u.Z}
}

func (v Vector) AddScalar(s float64) Vector {
	return Vector{v.X + s, v.Y + s, v.Z + s}
}

func (v Vector) SubtractVector(u Vector) Vector {
	return Vector{v.X - u.X, v.Y - u.Y, v.Z - u.Z}
}

func (v Vector) MultiplyVector(u Vector) Vector {
	return Vector{v.X * u.X, v.Y * u.Y, v.Z * u.Z}
}

func (v Vector) MultiplyScalar(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s}
}

func (v Vector) DivideScalar(s float64) Vector {
	return v.MultiplyScalar(1 / s)
}

func (v Vector) DotProduct(u Vector) float64 {
	return (v.X * u.X) + (v.Y * u.Y) + (v.Z * u.Z)
}

func (v Vector) MultiplyComponents(u Vector) Vector {
	return NewVector(v.X*u.X, v.Y*u.Y, v.Z*u.Z)
}

func (v Vector) CrossProduct(u Vector) Vector {
	x := (v.Y * u.Z) - (v.Z * u.Y)
	y := -((v.X * u.Z) - (v.Z * u.X))
	z := (v.X * u.Y) - (v.Y * u.X)

	return Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) UnitVector() Vector {
	return v.DivideScalar(v.Length())
}

func (v Vector) Reflect(n Vector) Vector {
	b := 2 * v.DotProduct(n)
	return v.SubtractVector(n.MultiplyScalar(b))
}

func (v Vector) Refract(o Vector, n float64) (bool, Vector) {
	uv := v.UnitVector()
	uo := o.UnitVector()
	dt := uv.DotProduct(uo)
	discriminant := 1.0 - (n * n * (1 - dt*dt))
	if discriminant > 0 {
		a := uv.SubtractVector(o.MultiplyScalar(dt)).MultiplyScalar(n)
		b := o.MultiplyScalar(math.Sqrt(discriminant))
		return true, a.SubtractVector(b)
	}
	return false, Vector{}
}
