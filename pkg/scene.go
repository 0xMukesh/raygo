package pkg

type Scene struct {
	Elements []Hittable
}

func (s Scene) Hit(r Ray, tMin, tMax float64) (bool, *HitRecord) {
	hitAnything := false
	closest := tMax
	record := &HitRecord{}

	for _, e := range s.Elements {
		hit, tempRecord := e.Hit(r, tMin, closest)

		if hit {
			hitAnything = true
			closest = tempRecord.T
			record = tempRecord
		}
	}

	return hitAnything, record
}
