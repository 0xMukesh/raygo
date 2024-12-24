package pkg

import (
	"math"
)

type Sphere struct {
	Centre Vector
	Radius float64
	Material
}

func NewSphere(centre Vector, radius float64, material Material) Sphere {
	return Sphere{
		Centre:   centre,
		Radius:   radius,
		Material: material,
	}
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	oc := r.Origin.SubtractVector(s.Centre)

	a := r.Direction.DotProduct(r.Direction)
	b := 2 * r.Direction.DotProduct(oc)
	c := oc.DotProduct(oc) - (s.Radius * s.Radius)

	discriminant := b*b - 4*a*c

	hitRecord := &HitRecord{Material: s.Material}

	if discriminant >= 0 {
		t := (-b - math.Sqrt(discriminant)) / (2 * a)

		if tMin <= t && tMax >= t {
			hitRecord.T = t
			hitRecord.P = r.At(t)
			hitRecord.N = hitRecord.P.SubtractVector(s.Centre).DivideScalar(s.Radius)
			return true, hitRecord
		}

		t = (-b + math.Sqrt(discriminant)) / (2 * a)

		if tMin <= t && tMax >= t {
			hitRecord.T = t
			hitRecord.P = r.At(t)
			hitRecord.N = hitRecord.P.SubtractVector(s.Centre).DivideScalar(s.Radius)
			return true, hitRecord
		}
	}

	return false, hitRecord
}
