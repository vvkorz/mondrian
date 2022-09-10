package main

import (
	"image/color"
	"mondrian/mondrian"
)

func main() {
	// Define canvas size
	x0 := 0
	x1 := 150
	y0 := 0
	y1 := 150
	// Complexity of the partitioning: the lower the number, the less partitions are made.
	//Do not input a number larger than 0.08 for the 150x150 canvas
	//If choosing the onlyRect option (draws only rectangles without the lines, go for
	//higher complexity, e.g. 0.05).
	cmplx := 0.03

	// Draws rectangles and lines
	mondrian.Draw(x0, x1, y0, y1, cmplx)

	// Draws only rectangles with higher complexity and given colors and color distribution
	var cols = []color.Color{
		color.White,
		color.RGBA{255, 0, 0, 255},
		color.RGBA{255, 255, 0, 255},
		color.RGBA{0, 0, 255, 255},
	}
	var probs = []float64{
		0.3,
		0.4,
		0.3,
		0.3,
	}
	mondrian.DrawR(x0, x1, y0, y1, 0.07, true, cols, probs)
}
