package pkg

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"
)

type Camera struct {
	LowerLeft, Origin, Horizontal, Vertical    Vector
	FocalLength, ViewportHeight, ViewportWidth float64
	MaxDepth                                   int
}

func NewCamera(lookFrom, lookAt, vUp Vector, vfov, aspectRatio float64) *Camera {
	c := new(Camera)

	c.Origin = lookFrom
	c.MaxDepth = 50
	c.FocalLength = lookFrom.SubtractVector(lookAt).Length()

	theta := vfov * math.Pi / 180
	h := math.Tan(theta / 2)

	c.ViewportHeight = 2 * h * c.FocalLength
	c.ViewportWidth = aspectRatio * c.ViewportHeight

	w := lookFrom.SubtractVector(lookAt).UnitVector()
	u := vUp.CrossProduct(w).UnitVector()
	v := w.CrossProduct(u)

	viewportU := u.MultiplyScalar(c.ViewportWidth)
	viewportV := v.MultiplyScalar(c.ViewportHeight)

	c.LowerLeft = c.Origin.SubtractVector(viewportU.DivideScalar(2)).SubtractVector(viewportV.DivideScalar(2)).SubtractVector(w.MultiplyScalar(c.FocalLength))
	c.Horizontal = u.MultiplyScalar(c.ViewportWidth)
	c.Vertical = v.MultiplyScalar(c.ViewportHeight)

	return c
}

func (c *Camera) RayAt(u, v float64) Ray {
	horizontal := c.Horizontal.MultiplyScalar(u)
	vertical := c.Vertical.MultiplyScalar(v)
	position := horizontal.AddVector(vertical)
	direction := c.LowerLeft.AddVector(position)

	return Ray{c.Origin, direction}
}

func (c *Camera) Render(scene Scene, imageHeight, imageWidth, numberOfSamples int, f *os.File) {
	for j := imageHeight; j >= 0; j-- {
		fmt.Printf("scanlines remaining: %d\n", j)
		for i := 0; i < int(imageWidth); i++ {
			rgb := Vector{}

			for s := 0; s < numberOfSamples; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth)
				v := (float64(j) + rand.Float64()) / float64(imageHeight)

				ray := c.RayAt(u, v)
				color := c.RayColor(ray, scene, 0)
				rgb = rgb.AddVector(color.ToVector())
			}

			rgb = rgb.DivideScalar(float64(numberOfSamples))

			if err := WriteColor(f, rgb.ToColor()); err != nil {
				panic(err.Error())
			}
		}
	}

	fmt.Println("done")
}

func (c *Camera) RayColor(r Ray, h Hittable, depth int) Color {
	if depth >= c.MaxDepth {
		return NewColor(0, 0, 0)
	}

	white := NewColor(1, 1, 1).ToVector()
	blue := NewColor(0.5, 0.7, 1).ToVector()

	// using 0.0001 as the initial value instead of 0 to avoid shadow acne caused due to floating number rounding issues
	found, rec := h.Hit(r, 0.0001, math.MaxFloat64)
	if found {
		found, ray := rec.Scatter(r, rec)
		if found {
			newColor := c.RayColor(ray, h, depth+1).ToVector()
			return rec.Material.Color().ToVector().MultiplyComponents(newColor).ToColor()
		}
	}

	unitVector := r.Direction.UnitVector()
	a := (1 + unitVector.Y) * 0.5
	// linear blend: blended_value = ((1 - a) * (start_value)) + (a * end_value)
	return white.MultiplyScalar(1 - a).AddVector(blue.MultiplyScalar(a)).ToColor()
}
