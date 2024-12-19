package pkg

import "math"

type Camera struct {
	LowerLeft, Origin, Horizontal, Vertical    Vector
	FocalLength, ViewportHeight, ViewportWidth float64
}

func NewCamera(aspectRatio float64) *Camera {
	c := new(Camera)

	c.FocalLength = 1.0
	c.ViewportHeight = 2.0
	c.ViewportWidth = c.ViewportHeight * aspectRatio

	c.LowerLeft = NewVector(-c.ViewportWidth/2.0, -c.ViewportHeight/2, -c.FocalLength)
	c.Horizontal = NewVector(c.ViewportWidth, 0, 0)
	c.Vertical = NewVector(0, c.ViewportHeight, 0)
	c.Origin = NewVector(0.0, 0.0, 0.0)

	return c
}

func (c *Camera) RayAt(u, v float64) Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return Ray{c.Origin, direction}
}

func (c *Camera) position(u, v float64) Vector {
	horizontal := c.Horizontal.MultiplyScalar(u)
	vertical := c.Vertical.MultiplyScalar(v)

	return horizontal.AddVector(vertical)
}

func (c *Camera) direction(position Vector) Vector {
	return c.LowerLeft.AddVector(position)
}

func (c *Camera) RayColor(r Ray, h Hittable) Color {
	white := NewColor(1, 1, 1).ToVector()
	blue := NewColor(0.5, 0.7, 1).ToVector()

	found, rec := h.Hit(r, 0, math.MaxFloat64)
	if found {
		return rec.N.AddScalar(1).MultiplyScalar(0.5).ToColor()
	}

	unitVector := r.Direction.UnitVector()
	a := (1 + unitVector.Y) * 0.5
	// linear blend: blended_value = ((1 - a) * (start_value)) + (a * end_value)
	return white.MultiplyScalar(1 - a).AddVector(blue.MultiplyScalar(a)).ToColor()
}
