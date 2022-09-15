package mondrian

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

// Draw calls DrawR without the additional parameters onlyRect, c and p
// and thus draws both the Rectangles and Lines on the same initial canvas
// under given complexity cmplx.
func Draw(
	x0 int,
	x1 int,
	y0 int,
	y1 int,
	cmplx float64,
	path string,
) *image.RGBA {
	return DrawR(x0, x1, y0, y1, cmplx, false, Colors, ColorProbs, path)
}

// DrawR takes a given canvas size and calls functions Rectangles and Lines to generate
// a single image given the partition complexity cmplx. If onlyRect is true, then the image
// is generated only with Rectangles. The colors of the partitions are set in cols and the
// distribution of colors in probs. Image (png) is stored to path.
func DrawR(
	x0 int,
	x1 int,
	y0 int,
	y1 int,
	cmplx float64,
	onlyRect bool,
	cols []color.Color,
	probs []float64,
	path string,
) *image.RGBA {
	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	// Calculating the partitioning budget
	lambda := float64(x1 - x0 + y1 - y0)

	// Initializing the Mondrian Process
	mp := MProcess{
		Lambda: lambda,
		X0:     x0,
		X1:     x1,
		Y0:     y0,
		Y1:     y1,
	}
	// Initializing the first rectangle in canvas size
	r := []image.Rectangle{
		{Min: image.Pt(mp.X0, mp.Y0), Max: image.Pt(mp.X1, mp.Y1)},
	}
	pr := Rectangles(mp, &r, true, cmplx)

	// Generating the empty canvas for the drawing of rectangles
	rectImage := image.NewRGBA((*pr)[0])

	// Drawing rectangles
	for _, rect := range (*pr)[1:] {
		c := RndColor(cols, probs)
		draw.Draw(rectImage, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	if onlyRect {
		f, err := os.Create(path)
		if err != nil {
			log.Fatalf("failed create file: %s", err)
		}
		png.Encode(f, rectImage)
		return rectImage
	}

	// Drawing lines
	lineImage := image.NewRGBA(image.Rectangle{image.Point{x0, y0}, image.Point{x1, y1}})

	Lines(pr, lineImage, color.RGBA{30, 38, 33, 255})

	// Overlaying the two images
	draw.Draw(rectImage, r[0], lineImage, image.ZP, draw.Over)

	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(f, rectImage)
	return rectImage
}
