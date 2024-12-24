package pkg

import (
	"math"
	"math/rand/v2"
)

type Material interface {
	Scatter(r Ray, h *HitRecord) (bool, Ray)
	Color() Color
}

type Lambertian struct {
	C Color
}

func NewLambertianMaterial(color Color) Lambertian {
	return Lambertian{color}
}

func (l Lambertian) Scatter(r Ray, h *HitRecord) (bool, Ray) {
	direction := RandomOnHemisphere(h.N).AddVector(RandomUnitVector())
	ray := NewRay(h.P, direction)
	return true, ray
}

func (l Lambertian) Color() Color {
	return l.C
}

type Metal struct {
	C    Color
	Fuzz float64
}

func NewMetalMaterial(color Color, fuzz float64) Metal {
	return Metal{color, fuzz}
}

func (m Metal) Scatter(r Ray, h *HitRecord) (bool, Ray) {
	reflectedRayDirection := r.Direction.Reflect(h.N).AddVector(RandomUnitVector().MultiplyScalar(m.Fuzz))
	reflectedRay := NewRay(h.P, reflectedRayDirection)
	isReflected := reflectedRayDirection.DotProduct(h.N) > 0

	return isReflected, reflectedRay
}

func (m Metal) Color() Color {
	return m.C
}

func (m Metal) Reflect(i Vector, n Vector, origin Vector) Vector {
	b := n.MultiplyScalar(-i.DotProduct(n))
	reflectedRayDirection := i.AddVector(b.MultiplyScalar(2))
	return reflectedRayDirection
}

type Dielectric struct {
	RefractiveIndex float64
}

func NewDielectricMaterial(refractiveIndex float64) Dielectric {
	return Dielectric{refractiveIndex}
}

func (d Dielectric) Color() Color {
	return NewColor(1, 1, 1)
}

func (d Dielectric) Scatter(r Ray, h *HitRecord) (bool, Ray) {
	var outwardNormal Vector
	var niOverNt, cosine float64

	if r.Direction.DotProduct(h.N) > 0 {
		outwardNormal = h.N.MultiplyScalar(-1)
		niOverNt = d.RefractiveIndex

		a := r.Direction.DotProduct(h.N) * d.RefractiveIndex
		b := r.Direction.Length()

		cosine = a / b
	} else {
		outwardNormal = h.N
		niOverNt = 1.0 / d.RefractiveIndex

		a := r.Direction.DotProduct(h.N) * d.RefractiveIndex
		b := r.Direction.Length()

		cosine = -a / b
	}

	var success bool
	var refracted Vector
	var reflectProbability float64

	if success, refracted = r.Direction.Refract(outwardNormal, niOverNt); success {
		reflectProbability = d.Schlick(cosine)
	} else {
		reflectProbability = 1.0
	}

	if rand.Float64() < reflectProbability {
		reflected := r.Direction.Reflect(h.N)
		return true, Ray{h.P, reflected}
	}

	return true, Ray{h.P, refracted}
}

func (d Dielectric) Schlick(cosine float64) float64 {
	r0 := (1 - d.RefractiveIndex) / (1 + d.RefractiveIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
