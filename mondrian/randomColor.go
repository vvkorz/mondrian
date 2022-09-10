package mondrian

import (
	"image/color"
	"log"
	"math/rand"
)

// RandomColor picks a randomly chosen color from either mainCol (the primary color
// to be picked with the highest probability) or the slice secondCols (secondary colors
// to be picked uniformly).
func RandomColor(mainCol color.Color, secondCols []color.Color) color.Color {
	if len(secondCols) != 3 {
		log.Fatalf("Length of secondCols has to be 3, input has %d\n",
			len(secondCols))
	}
	// Choosing if the color will be white or primary
	if rand.Float64() > 0.7 {
		red := 0.4
		yellow := 0.8

		switch f := rand.Float64(); {
		case f < red:
			// red
			return secondCols[0]
		case f >= red && f < yellow:
			// yellow
			return secondCols[1]
		case f >= yellow:
			// blue
			return secondCols[2]
		}
		return secondCols[rand.Intn(len(secondCols))]
	} else {
		return mainCol
	}
}
