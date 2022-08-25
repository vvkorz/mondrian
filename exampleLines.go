package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func exampleLines() {
	fmt.Println("drawing lines")
	width := 200
	height := 300
	wlines := []int{20, 130}
	hlines := []int{123, 222}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for i, s := range wlines {
		fmt.Println(i, s)
		for y := 0; y < height; y++ {
			img.Set(s, y, color.Black)
		}
	}
	for i, s := range hlines {
		fmt.Println(i, s)
		for x := 0; x < width; x++ {
			img.Set(x, s, color.Black)
		}
	}

	// Encode as PNG.
	f, _ := os.Create("img/exampleLines.png")
	png.Encode(f, img)
}
