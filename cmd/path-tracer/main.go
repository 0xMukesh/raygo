package main

import (
	"fmt"
	"math/rand/v2"
	"os"

	"github.com/0xmukesh/path-tracer/pkg"
)

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	camera := pkg.NewCamera(aspectRatio)

	f, err := os.Create("test.ppm")
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	if err != nil {
		panic(err.Error())
	}

	ns := 10

	for j := imageHeight; j >= 0; j-- {
		fmt.Printf("scanlines remaining: %d\n", j)

		for i := 0; i < int(imageWidth); i++ {
			rgb := pkg.Vector{}

			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth)
				v := (float64(j) + rand.Float64()) / float64(imageHeight)

				ray := camera.RayAt(u, v)
				color := ray.RayColor()
				rgb = rgb.AddVector(color.ToVector())
			}

			rgb = rgb.DivideScalar(float64(ns))

			if err := pkg.WriteColor(f, rgb.ToColor()); err != nil {
				panic(err.Error())
			}
		}
	}

	fmt.Println("done")
}
