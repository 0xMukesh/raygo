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

	lookFrom := pkg.NewVector(1, 2, 1)
	lookAt := pkg.NewVector(0, 0, -1)
	vUp := pkg.NewVector(0, 1, 0)

	camera := pkg.NewCamera(lookFrom, lookAt, vUp, 30, float64(imageWidth)/float64(imageHeight))

	f, err := os.Create("test.ppm")
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	if err != nil {
		panic(err.Error())
	}

	ns := 300

	materialGround := pkg.NewLambertianMaterial(pkg.NewColor(0.8, 0.8, 0))
	materialLambertian := pkg.NewLambertianMaterial(pkg.NewColor(0.1, 0.2, 0.5))
	materialGlass := pkg.NewDielectricMaterial(1.5)
	materialBubble := pkg.NewDielectricMaterial(1.0 / 1.5)
	materialMetal := pkg.NewMetalMaterial(pkg.NewColor(0.8, 0.6, 0.2), 0)

	ground := pkg.NewSphere(pkg.NewVector(0, -100.5, -1), 100, materialGround)
	centre := pkg.NewSphere(pkg.NewVector(0, 0, -1.2), 0.5, materialLambertian)
	left := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.5, materialGlass)
	bubble := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.4, materialBubble)
	right := pkg.NewSphere(pkg.NewVector(1, 0, -1), 0.5, materialMetal)
	topLeft := pkg.NewSphere(pkg.NewVector(-1, 1, -1), 0.5, materialGlass)
	topMiddle := pkg.NewSphere(pkg.NewVector(0, 1, -1), 0.5, materialMetal)
	topRight := pkg.NewSphere(pkg.NewVector(1, 1, -1), 0.5, materialGlass)

	scene := pkg.Scene{Elements: []pkg.Hittable{&ground, &centre, &left, &bubble, &right, &topLeft, &topMiddle, &topRight}}

	camera.Render(scene, imageHeight, imageWidth, ns, f)
}
