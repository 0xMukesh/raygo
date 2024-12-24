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

	lookFrom := pkg.NewVector(0.2, 0.01, 1.5)
	lookAt := pkg.NewVector(0, 0.5, -1)
	vUp := pkg.NewVector(0, 1, 0)

	camera := pkg.NewCamera(lookFrom, lookAt, vUp, 40, float64(imageWidth)/float64(imageHeight))

	f, err := os.Create("test.ppm")
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	if err != nil {
		panic(err.Error())
	}

	ns := 200

	materialGround := pkg.NewLambertianMaterial(pkg.NewColor(0.8, 0.8, 0))
	materialCentre := pkg.NewLambertianMaterial(pkg.NewColor(0.1, 0.2, 0.5))
	materialLeft := pkg.NewDielectricMaterial(1.5)
	materialBubble := pkg.NewDielectricMaterial(1.0 / 1.5)
	materialRight := pkg.NewMetalMaterial(pkg.NewColor(0.8, 0.6, 0.2), 0)

	ground := pkg.NewSphere(pkg.NewVector(0, -100.5, -1), 100, materialGround)
	centre := pkg.NewSphere(pkg.NewVector(0, 0, -1.2), 0.5, materialCentre)
	left := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.5, materialLeft)
	bubble := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.4, materialBubble)
	right := pkg.NewSphere(pkg.NewVector(1, 0, -1), 0.5, materialRight)

	scene := pkg.Scene{Elements: []pkg.Hittable{&ground, &centre, &left, &bubble, &right}}

	camera.Render(scene, imageHeight, imageWidth, ns, f)
}
