package pkg

import "math"

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

func (r Ray) HitsSphere(centre Vector, radius float64) bool {
	oc := centre.SubtractVector(r.Origin)

	a := r.Direction.DotProduct(r.Direction)
	b := -2 * r.Direction.DotProduct(oc)
	c := oc.DotProduct(oc) - math.Pow(radius, 2)

	return math.Pow(b, 2)-4*a*c >= 0
}

func RayColor(r Ray) Color {
	if r.HitsSphere(NewVector(0, 0, -1), 0.5) {
		return NewVector(1, 0, 0)
	}

	unitVector := r.Direction.UnitVector()
	a := (1 - unitVector.X) * 0.5
	// lerp - (1 - a) * (start value) + (a) * (end value)
	return NewVector(1, 1, 1).MultiplyScalar(1 - a).AddVector(NewVector(0.5, 0.7, 1).MultiplyScalar(a))
}
