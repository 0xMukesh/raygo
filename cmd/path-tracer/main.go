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

	materialGround := pkg.NewLambertianMaterial(pkg.NewColor(0.8, 0.8, 0))
	materialSphere := pkg.NewLambertianMaterial(pkg.NewColor(0.1, 0.2, 0.5))
	materialMetal := pkg.NewMetalMaterial(pkg.NewColor(0.8, 0.6, 0.2), 0.3)
	materialGlass := pkg.NewDielectricMaterial(1.5)
	materialBubble := pkg.NewDielectricMaterial(1 / 1.5)

	ground := pkg.NewSphere(pkg.NewVector(0, -100.5, -1), 100, materialGround)
	sphere := pkg.NewSphere(pkg.NewVector(0, 0, -1), 0.5, materialSphere)
	metal := pkg.NewSphere(pkg.NewVector(1, 0, -1), 0.5, materialMetal)
	glass := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.5, materialGlass)
	bubble := pkg.NewSphere(pkg.NewVector(-1, 0, -1), 0.4, materialBubble)

	scene := pkg.Scene{Elements: []pkg.Hittable{&ground, &sphere, &metal, &glass, &bubble}}

	camera.Render(scene, imageHeight, imageWidth, ns, f)
}
