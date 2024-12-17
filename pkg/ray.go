package pkg

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

func RayColor(r Ray) Color {
	unitVector := r.Direction.UnitVector()
	a := 0.5 * (-unitVector.X + 1)
	return NewVector(1, 1, 1).MultiplyScalar(1 - a).AddVector(NewVector(0.5, 0.7, 1).MultiplyScalar(a))
}
