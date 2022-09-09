package main

import (
	"image"
	"image/color"
)

// hLine draws a horizontal line pixel by pixel from x0 to x1 at point y
func hLine(x0 int, x1 int, y int, img *image.RGBA, col color.Color) {
	for ; x0 <= x1; x0++ {
		img.Set(x0, y, col)
	}
}

// vLine draws a vertical line pixel by pixel from y0 to y1 at point x
func vLine(x int, y0 int, y1 int, img *image.RGBA, col color.Color) {
	for ; y0 <= y1; y0++ {
		img.Set(x, y0, col)
	}
}

// drawLines draws horizontal and vertical lines as defined by the rectangle edges
//of rectangles stored at pr, in color col into img.
func drawLines(pr *[]image.Rectangle, img *image.RGBA, col color.Color) {
	for _, r := range (*pr)[1:] {
		// Drawing horizontal lines
		hLine(r.Min.X, r.Max.X, r.Min.Y, img, col)
		hLine(r.Min.X, r.Max.X, r.Max.Y, img, col)
		// Drawing vertical lines
		vLine(r.Min.X, r.Min.Y, r.Max.Y, img, col)
		vLine(r.Max.X, r.Min.Y, r.Max.Y, img, col)
	}
}
