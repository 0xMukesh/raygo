package main

import (
	"fmt"
	"os"

	"github.com/0xmukesh/path-tracer/pkg"
)

func main() {
	imageWidth := 400
	aspectRatio := 16.0 / 9.0
	imageHeight := int(float64(imageWidth) / aspectRatio)

	focalLength := 1.0
	viewportHeight := 2.0
	viewportWidth := viewportHeight * (float64(imageWidth) / float64(imageHeight))
	cameraCentre := pkg.NewVector(0, 0, 0)

	viewportU := pkg.NewVector(viewportWidth, 0, 0)
	viewportV := pkg.NewVector(0, -viewportHeight, 0)

	pixelDeltaU := viewportU.DivideScalar(float64(imageWidth))
	pixelDeltaV := viewportV.DivideScalar(float64(imageHeight))

	focalLengtVector := pkg.NewVector(0, 0, focalLength)
	viewportUpperLeft := focalLengtVector.SubtractVector(cameraCentre).SubtractVector(viewportU.DivideScalar(2)).SubtractVector(viewportV.DivideScalar(2))
	pixel00Loc := viewportUpperLeft.AddVector(pixelDeltaU.DivideScalar(2)).AddVector(pixelDeltaV.DivideScalar(2))

	f, err := os.Create("test.ppm")
	if err != nil {
		panic(err.Error())
	}

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", imageWidth, imageHeight)
	if err != nil {
		panic(err.Error())
	}

	for j := 0; j < imageHeight; j++ {
		fmt.Printf("scanlines remaining: %d\n", imageHeight-j)
		for i := 0; i < imageWidth; i++ {
			pixelCentre := pixel00Loc.AddVector(pixelDeltaU.MultiplyScalar(float64(i))).AddVector(pixelDeltaV.MultiplyScalar(float64(j)))
			rayDirection := pixelCentre.SubtractVector(cameraCentre)

			ray := pkg.NewRay(cameraCentre, rayDirection)
			pixelColor := pkg.RayColor(ray)

			if err := pkg.WriteColor(f, pixelColor); err != nil {
				panic(err.Error())
			}
		}
	}

	fmt.Println("done")
}
