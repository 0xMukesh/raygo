package pkg

import (
	"math"
	"math/rand/v2"
)

// returns a random number in [min, max) range
func RandomInRange(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

// returns a vector with random components
func RandomVectorInRange(min, max float64) Vector {
	return NewVector(RandomInRange(min, max), RandomInRange(min, max), RandomInRange(min, max))
}

// returns a random unit vector which is present inside the unit sphere
func RandomUnitVector() Vector {
	for {
		v := NewVector(RandomInRange(-1, 1), RandomInRange(-1, 1), RandomInRange(-1, 1))
		if math.MinInt64 < v.Length()*v.Length() && v.Length()*v.Length() <= 1 {
			return v.UnitVector()
		}
	}
}

// returns a random unit vector which is present on the hemisphere (by taking dot product of normal vector at that point and the random unit vector)
func RandomOnHemisphere(normal Vector) Vector {
	onUnitSphere := RandomUnitVector()

	if onUnitSphere.DotProduct(normal) > 0 {
		return onUnitSphere
	} else {
		return onUnitSphere.MultiplyScalar(-1)
	}
}

func LinearToGamma(linear float64) float64 {
	if linear > 0 {
		return math.Sqrt(linear)
	} else {
		return 0
	}
}
