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

func NewCamera(aspectRatio float64) *Camera {
	c := new(Camera)

	c.FocalLength = 1.0
	c.ViewportHeight = 2.0
	c.ViewportWidth = c.ViewportHeight * aspectRatio
	c.MaxDepth = 50

	c.LowerLeft = NewVector(-c.ViewportWidth/2.0, -c.ViewportHeight/2, -c.FocalLength)
	c.Horizontal = NewVector(c.ViewportWidth, 0, 0)
	c.Vertical = NewVector(0, c.ViewportHeight, 0)
	c.Origin = NewVector(0.0, 0.0, 0.0)

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
				color := c.RayColor(ray, scene, c.MaxDepth)
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
	if depth <= 0 {
		return NewColor(0, 0, 0)
	}

	white := NewColor(1, 1, 1).ToVector()
	blue := NewColor(0.5, 0.7, 1).ToVector()

	// using 0.0001 as the initial value instead of 0 to avoid shadow acne caused due to floating number rounding issues
	found, rec := h.Hit(r, 0.0001, math.MaxFloat64)
	if found {
		found, ray := rec.Scatter(r, rec)
		if found {
			newColor := c.RayColor(ray, h, depth-1).ToVector()
			return rec.Material.Color().MultiplyComponents(newColor).ToColor()
		}
	}

	unitVector := r.Direction.UnitVector()
	a := (1 + unitVector.Y) * 0.5
	// linear blend: blended_value = ((1 - a) * (start_value)) + (a * end_value)
	return white.MultiplyScalar(1 - a).AddVector(blue.MultiplyScalar(a)).ToColor()
}
