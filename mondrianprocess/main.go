// package mondrianprocess
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func randomColor(mainCol color.Color, secondCols []color.Color) color.Color {
	if len(secondCols) != 3 {
		log.Fatalf("Length of secondCols has to be 3, input has %d\n", len(secondCols))
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

func main() {
	x0 := 0
	x1 := 150
	y0 := 0
	y1 := 150
	lambda := float64(x1 - x0 + y1 - y0)
	// Initializing the Mondrian Process
	mp := mProcess{
		lambda: lambda,
		x0:     x0,
		x1:     x1,
		y0:     y0,
		y1:     y1,
	}
	// Initializing the Rectangles slice
	r := []image.Rectangle{
		{Min: image.Pt(mp.x0, mp.y0), Max: image.Pt(mp.x1, mp.y1)},
	}

	pr := mondrianProcess(mp, &r)

	rectImage := image.NewRGBA((*pr)[0])

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	for inx, rect := range (*pr)[1:] {
		fmt.Println(len(colors[1:]))
		c := randomColor(colors[0], colors[1:])
		fmt.Println(inx, rect, c)
		draw.Draw(rectImage, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	file, err := os.Create("exampleRectangles.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)
}
