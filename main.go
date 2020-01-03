package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	// color for the rect
	red := color.RGBA{255, 0, 0, 1}

	img := gocv.IMRead("InputImage.jpg", gocv.IMReadAnyColor)
	template := gocv.IMRead("TemplateImage.jpg", gocv.IMReadAnyColor)

	matResult := gocv.NewMat()
	mask := gocv.NewMat()

	// TmCcoeffNormed - works
	// TmCcorrNormed - works
	// TmSqdiffNormed - doesn't work
	gocv.MatchTemplate(img, template, &matResult, gocv.TmCcorrNormed, mask)
	defer mask.Close()

	minConfidence, maxConfidence, minLoc, maxLoc := gocv.MinMaxLoc(matResult)
	defer matResult.Close()
	fmt.Println("minConfidence:", minConfidence)
	fmt.Println("maxConfidence:", maxConfidence)
	fmt.Println("minLoc:", minLoc)
	fmt.Println("maxLoc:", maxLoc)

	dims := template.Size()
	// dims[1] = width
	// dims[0] = height

	// can also use Rows & Cols for (x,y)
	// template.Cols() is width
	// template.Rows() is height

	//fmt.Println("Template Size: ", dims[1], dims[0]) // width & height
	//fmt.Println("Template Col/Rows: ", template.Cols(), template.Rows())

	point := image.Point{maxLoc.X + dims[1], maxLoc.Y + dims[0]}

	gocv.Rectangle(&img, image.Rectangle{Min: maxLoc, Max: point}, red, 1)

	gocv.IMWrite("OutputImage.jpg", img)

	/*
		window := gocv.NewWindow("Template Matching")

		for {
			window.IMShow(img)
			window.WaitKey(10)
		}
	*/

}
