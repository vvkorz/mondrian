package main

import (
	"flag"
	"fmt"
	"image/color"
	"mondrianart/mondrian"
	"path/filepath"
)

func main() {
	// Cmd arguments where to store the resulting image
	imgpath := flag.String(
		"imgpath",
		"./img",
		"Path to resulting images, e.g. ~/mylocalpath",
	)
	flag.Parse()
	// Define canvas size
	x0 := 0
	x1 := 500
	y0 := 0
	y1 := 500
	// Complexity of the partitioning: the lower the number, the less partitions are made.
	//Do not input a number larger than 0.05 for the 500x500 canvas
	//If choosing the onlyRect option (draws only rectangles without the lines, go for
	//higher complexity, e.g. 0.03).
	cmplx := 0.01
	fmt.Printf("Saving your Mondrian images to %s\n", *imgpath)

	// Draws rectangles and lines
	mondrian.Draw(x0, x1, y0, y1, cmplx, filepath.Join(*imgpath, "RectangLines.png"))

	// Draws only rectangles with higher complexity and given colors and color distribution
	var cols = []color.Color{
		color.White,
		color.RGBA{221, 1, 0, 255},
		color.RGBA{250, 201, 1, 255},
		color.RGBA{34, 80, 149, 255},
		color.RGBA{30, 38, 33, 255},
	}
	var probs = []float64{
		0.5,
		0.13,
		0.13,
		0.13,
		0.11,
	}
	mondrian.DrawR(x0, x1, y0, y1, 0.03, true, cols, probs, filepath.Join(*imgpath, "Rectangles.png"))
}
