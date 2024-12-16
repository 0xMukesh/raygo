package main

import (
	"fmt"
	"os"
)

func main() {
	width := 256
	height := 400

	f, err := os.Create("test.ppm")
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", width, height)
	if err != nil {
		panic(err.Error())
	}

	color := 255.999

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r := float64(i) / float64(width)
			g := float64(j) / float64(height)
			b := 0.0

			ir := int(color * r)
			ig := int(color * g)
			ib := int(color * b)

			_, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			if err != nil {
				panic(err.Error())
			}
		}
	}

}
