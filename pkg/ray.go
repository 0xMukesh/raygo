package pkg

import (
	"math"
)

type Ray struct {
	Origin, Direction Vector
}

func NewRay(origin, direction Vector) Ray {
	return Ray{
		Origin:    origin,
		Direction: direction,
	}
}

func (r Ray) At(t float64) Vector {
	return r.Origin.AddVector(r.Direction.MultiplyScalar(t))
}

func (r Ray) HitsSphere(centre Vector, radius float64) (bool, float64) {
	oc := centre.SubtractVector(r.Origin)

	a := r.Direction.DotProduct(r.Direction)
	b := -2 * r.Direction.DotProduct(oc)
	c := oc.DotProduct(oc) - math.Pow(radius, 2)

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant >= 0 {
		return true, (-b - math.Sqrt(discriminant)) / (2 * a)
	} else {
		return false, 0
	}
}

func RayColor(r Ray) Color {
	centre := NewVector(0, 0, -1)
	radius := 0.5

	found, t := r.HitsSphere(centre, radius)
	if found {
		normal := (r.At(t).SubtractVector(centre)).DivideScalar(radius)
		return NewVector(normal.X+1, normal.Y+1, normal.Z+1).MultiplyScalar(0.5).ToColor()
	}

	unitVector := r.Direction.UnitVector()
	a := (1 + unitVector.Y) * 0.5
	// lerp - (1 - a) * (start value) + (a) * (end value)
	return NewVector(1, 1, 1).MultiplyScalar(1 - a).AddVector(NewVector(0.5, 0.7, 1).MultiplyScalar(a)).ToColor()
}
