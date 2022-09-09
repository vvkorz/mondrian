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

var colors = []color.Color{
	color.White,
	// red
	color.RGBA{255, 0, 0, 255},
	// yellow
	color.RGBA{255, 255, 0, 255},
	// blue
	color.RGBA{0, 0, 255, 255},
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

	pr := mondrianProcess(mp, &r, true, 0.04)

	rectImage := image.NewRGBA((*pr)[0])

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	for inx, rect := range (*pr)[1:] {
		c := randomColor(colors[0], colors[1:])
		fmt.Println(inx, rect, c)
		draw.Draw(rectImage, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	file, err := os.Create("exampleRectangles.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)

	// ________________ Drawing lines ________________
	lineImage := image.NewRGBA(image.Rectangle{image.Point{x0, y0}, image.Point{x1, y1}})

	drawLines(pr, lineImage, color.Black)

	// Encode as PNG.
	f, err := os.Create("exampleLines.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(f, lineImage)

	// Overlaying the two images
	//draw.DrawMask(rectImage, r[0], lineImage, image.ZP, mask, image.ZP, draw.Src)
	//dst Image, r image.Rectangle, src image.Image, sp image.Point, mask image.Image, mp image.Point, op Op
	draw.Draw(rectImage, r[0], lineImage, image.ZP, draw.Over)

	f, err = os.Create("drawover.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(f, rectImage)
}
