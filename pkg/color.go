package pkg

import (
	"fmt"
	"os"
)

func WriteColor(file *os.File, pixel Vector) error {
	r := pixel.X
	g := pixel.Y
	b := pixel.Z

	ir := int(255.999 * r)
	ig := int(255.999 * g)
	ib := int(255.999 * b)

	_, err := fmt.Fprintf(file, "%d %d %d\n", ir, ig, ib)
	return err
}
