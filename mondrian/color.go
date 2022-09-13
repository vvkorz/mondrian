package mondrian

import (
	"image/color"
	"log"
	"math/rand"
)

// Colors list allowed colors to be used to color the resulting partitions of Rectangles.
var Colors = []color.Color{
	color.White,
	// red
	color.RGBA{255, 0, 0, 255},
	// yellow
	color.RGBA{255, 255, 0, 255},
	// blue
	color.RGBA{0, 0, 255, 255},
}

// ColorProbs define the probabilities for Colors with the same order of index:
// white, red, yellow and blue.
var ColorProbs = []float64{
	0.6,
	0.4,
	0.3,
	0.3,
}

// RndColor picks a randomly chosen color from the slice of colors c taking the first indices
// as colors with the highest probability p of occurring.
func RndColor(
	c []color.Color,
	p []float64,
) color.Color {
	if len(c) != 4 {
		log.Fatalf("Length of c has to be 4, input has %d\n", len(c))
	}
	if p[1]+p[2]+p[3] != 1 {
		log.Fatalln("The sum of the last three probabilities has to be 1")
	}
	// Choosing if the color will be white or primary
	if rand.Float64() > p[0] {
		switch f := rand.Float64(); {
		case f < p[1]:
			// red
			return c[1]
		case f >= p[1] && f < (p[1]+p[2]):
			// yellow
			return c[2]
		case f >= (p[1] + p[2]):
			// blue
			return c[3]
		}
	}
	return c[0]
}
