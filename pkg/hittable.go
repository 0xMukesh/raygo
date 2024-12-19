package pkg

type HitRecord struct {
	N, P        Vector
	T           float64
	IsFrontFace bool
}

func (r *HitRecord) SetFaceNormal(ray Ray, outwardNormal Vector) {
	isFrontFace := ray.Direction.DotProduct(outwardNormal) < 0
	if isFrontFace {
		r.N = outwardNormal
	} else {
		r.N = outwardNormal.MultiplyScalar(-1)
	}
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool, *HitRecord)
}
