package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"math/rand"
	"mondrian/mondrian"
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
	mp := mondrian.MProcess{
		Lambda: lambda,
		X0:     x0,
		X1:     x1,
		Y0:     y0,
		Y1:     y1,
	}
	// Initializing the Rectangles slice
	r := []image.Rectangle{
		{Min: image.Pt(mp.X0, mp.Y0), Max: image.Pt(mp.X1, mp.Y1)},
	}

	pr := mondrian.MondrianProcess(mp, &r, true, 0.04)

	rectImage := image.NewRGBA((*pr)[0])

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	// Drawing rectangle
	for inx, rect := range (*pr)[1:] {
		c := mondrian.RandomColor(colors[0], colors[1:])
		fmt.Println(inx, rect, c)
		draw.Draw(rectImage, rect, &image.Uniform{c}, image.ZP, draw.Src)
	}

	file, err := os.Create("img/Rectangles.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)

	// Drawing lines
	lineImage := image.NewRGBA(image.Rectangle{image.Point{x0, y0}, image.Point{x1, y1}})

	mondrian.DrawLines(pr, lineImage, color.Black)

	// Encode as PNG.
	f, err := os.Create("img/Lines.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(f, lineImage)

	// Overlaying the two images
	draw.Draw(rectImage, r[0], lineImage, image.ZP, draw.Over)

	f, err = os.Create("img/Overlay.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(f, rectImage)
}
