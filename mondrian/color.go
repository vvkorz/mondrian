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
	color.RGBA{221, 1, 0, 255},
	// yellow
	color.RGBA{250, 201, 1, 255},
	// blue
	color.RGBA{34, 80, 149, 255},
	// black
	color.RGBA{30, 38, 33, 255},
}

// ColorProbs define the probabilities for Colors with the same order of index:
// white, red, yellow and blue.
var ColorProbs = []float64{
	0.5,
	0.13,
	0.13,
	0.13,
	0.11,
}

// RndColor picks a randomly chosen color from the slice of colors c taking the first indices
// as colors with the highest probability p of occurring.
func RndColor(
	c []color.Color,
	p []float64,
) color.Color {
	if len(c) != len(p) {
		log.Fatalf("Length of c %d != %d\n", len(c), len(p))
	}
	// calculate cumulative probabilities [0.4, 0.6] -> [0.4, 1]
	total_probability := 0.0
	random_float := rand.Float64() // choosing color
	return_indx := 0               // by default return first color
	for indx, probability := range p {
		total_probability += probability
		if random_float < total_probability {
			return c[indx]
		}
	}
	if total_probability != 1 {
		log.Fatalf("The sum of probabilities has to be 1 not %f\n", total_probability)
	}
	return c[return_indx]
}
