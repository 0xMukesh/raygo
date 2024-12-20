package pkg

type Material interface {
	Scatter(r Ray, h *HitRecord) (bool, Ray)
	Color() Vector
}

type Lambertian struct {
	C Vector
}

func NewLambertianMaterial(color Vector) Lambertian {
	return Lambertian{
		C: color,
	}
}

func (l Lambertian) Scatter(r Ray, h *HitRecord) (bool, Ray) {
	direction := RandomOnHemisphere(h.N).AddVector(RandomUnitVector())
	ray := NewRay(h.P, direction)
	return true, ray
}

func (l Lambertian) Color() Vector {
	return l.C
}

type Metal struct {
	C    Vector
	Fuzz float64
}

func NewMetalMaterial(color Vector, fuzz float64) Metal {
	return Metal{
		C:    color,
		Fuzz: fuzz,
	}
}

func (m Metal) Scatter(r Ray, h *HitRecord) (bool, Ray) {
	// reflected ray = v + 2b
	// b = -(v.n) * n

	b := h.N.MultiplyScalar(-r.Direction.DotProduct(h.N))
	reflectedRayDirection := r.Direction.AddVector(b.MultiplyScalar(2))
	reflectedRay := NewRay(h.P, reflectedRayDirection.AddVector(RandomOnHemisphere(h.N).MultiplyScalar(m.Fuzz)))
	isReflected := reflectedRayDirection.DotProduct(h.N) > 0

	return isReflected, reflectedRay
}

func (m Metal) Color() Vector {
	return m.C
}
