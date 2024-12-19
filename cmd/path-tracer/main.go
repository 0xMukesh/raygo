package main

import (
	"fmt"
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

	ns := 100

	sphere := pkg.NewSphere(pkg.NewVector(0, 0, -1), 0.5)
	floor := pkg.NewSphere(pkg.NewVector(0, -100.5, -1), 100)
	scene := pkg.Scene{Elements: []pkg.Hittable{&sphere, &floor}}

	camera.Render(scene, imageHeight, imageWidth, ns, f)
}
