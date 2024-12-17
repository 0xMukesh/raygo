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

func WriteColor(file *os.File, color Color) error {
	r := color.R
	g := color.G
	b := color.B

	ir := int(255.999 * r)
	ig := int(255.999 * g)
	ib := int(255.999 * b)

	_, err := fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
	return err
}
