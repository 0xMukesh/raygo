package pkg

type HitRecord struct {
	N, P Vector
	T    float64
	Material
}

type Hittable interface {
	Hit(r Ray, tMin, tMax float64) (bool, *HitRecord)
}
