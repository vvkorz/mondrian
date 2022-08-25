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

// The standard colors for a Mondrian image are red, yellow and black.
var colors = []color.Color{
	color.White,
	color.Black,
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
}

func exampleRectangles() {
	fmt.Println("drawing rectangles")
	// func Rect(x0, y0, x1, y1 int) Rectangle
	rectImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

	rect1 := image.Rect(0, 0, 120, 200)
	rect2 := image.Rect(120, 0, 200, 100)
	rect3 := image.Rect(120, 100, 200, 200)

	rectangles := [3]image.Rectangle{rect1, rect2, rect3}

	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	for inx, rect := range rectangles {
		fmt.Println(inx, rect)
		draw.Draw(rectImage, rect, &image.Uniform{colors[rand.Intn(len(colors))]}, image.ZP, draw.Src)
	}

	file, err := os.Create("img/exampleRectangles.png")
	if err != nil {
		log.Fatalf("failed create file: %s", err)
	}
	png.Encode(file, rectImage)
}
