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
