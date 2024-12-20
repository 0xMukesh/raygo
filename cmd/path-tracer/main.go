package main

import (
	"fmt"
	"os"

	"github.com/0xmukesh/path-tracer/pkg"
)

func main() {
	aspectRatio := 16.0 / 9.0
	imageWidth := 1000
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

	materialGround := pkg.NewLambertianMaterial(pkg.NewVector(0.8, 0.8, 0))
	materialCentre := pkg.NewLambertianMaterial(pkg.NewVector(0.1, 0.2, 0.5))
	materialLeft := pkg.NewMetalMaterial(pkg.NewVector(0.8, 0.8, 0.8), 0)
	materialRight := pkg.NewMetalMaterial(pkg.NewVector(0.8, 0.6, 0.2), 0)
	materialUp := pkg.NewMetalMaterial(pkg.NewVector(0.8, 0.2, 0.6), 0)

	ground := pkg.NewSphere(pkg.NewVector(0, -100.5, -1), 100, materialGround)
	centre := pkg.NewSphere(pkg.NewVector(0, 0, -1.2), 0.5, materialCentre)
	left := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.5, materialLeft)
	right := pkg.NewSphere(pkg.NewVector(1, 0, -1), 0.5, materialRight)
	up := pkg.NewSphere(pkg.NewVector(0, 1, -1.2), 0.5, materialUp)

	scene := pkg.Scene{Elements: []pkg.Hittable{&ground, &centre, &left, &right, &up}}

	camera.Render(scene, imageHeight, imageWidth, ns, f)
}
