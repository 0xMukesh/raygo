package pkg

type HitRecord struct {
	N, P        Vector
	T           float64
	IsFrontFace bool
	Material
}

func (h *HitRecord) SetFaceNormal(ray Ray, outwardNormal Vector) {
	isFrontFace := ray.Direction.DotProduct(outwardNormal) < 0
	if isFrontFace {
		h.N = outwardNormal
	} else {
		h.N = outwardNormal.MultiplyScalar(-1)
	}
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool, *HitRecord)
}
