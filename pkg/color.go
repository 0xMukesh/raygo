package pkg

import (
	"fmt"
	"os"
)

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{
		R: r,
		G: g,
		B: b,
	}
}

func (c Color) ToVector() Vector {
	return NewVector(c.R, c.G, c.B)
}

func WriteColor(file *os.File, color Color) error {
	r := LinearToGamma(color.R)
	g := LinearToGamma(color.G)
	b := LinearToGamma(color.B)

	ir := int(255.99 * r)
	ig := int(255.99 * g)
	ib := int(255.99 * b)

	_, err := fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
	return err
}
