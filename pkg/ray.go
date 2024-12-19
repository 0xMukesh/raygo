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

func (r Ray) RayColor() Color {
	sphereCentre := NewVector(0, 0, -1)
	sphereRadius := 0.5
	floorCentre := NewVector(0, -100.5, -1)
	floorRadius := 100.0

	white := NewColor(1, 1, 1).ToVector()
	blue := NewColor(0.5, 0.7, 1).ToVector()

	sphere := NewSphere(sphereCentre, sphereRadius)
	floor := NewSphere(floorCentre, floorRadius)

	scene := Scene{[]Hittable{&sphere, &floor}}

	found, rec := scene.Hit(r, 0, math.MaxFloat64)
	if found {
		return rec.N.AddScalar(1).MultiplyScalar(0.5).ToColor()
	}

	unitVector := r.Direction.UnitVector()
	a := (1 + unitVector.Y) * 0.5
	// linear blend: blended_value = ((1 - a) * (start_value)) + (a * end_value)
	return white.MultiplyScalar(1 - a).AddVector(blue.MultiplyScalar(a)).ToColor()
}
