package main

import (
	"fmt"
	"image/color"
	"mondrian/mondrian"
	"net/http"
	"strconv"
)

// handler reads and parses the html template for the first page
func handler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

// imageHandler responds to an incoming request by generating a Mondrian image fo size 500x500
// and according to the chosen style and complexity which is taken from the request form. The
// response is a html block with the image source and dimensions.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	x0 := 0
	x1 := 500
	y0 := 0
	y1 := 500
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cmplx, _ := strconv.ParseFloat(r.FormValue("complexity"), 64)
	style := r.FormValue("style")

	if style == "rectandline" {
		mondrian.Draw(x0, x1, y0, y1, cmplx, "img/mondrian_image.png")
	} else if style == "rect" {
		var cols = []color.Color{
			color.White,
			color.RGBA{255, 0, 0, 255},
			color.RGBA{255, 255, 0, 255},
			color.RGBA{0, 0, 255, 255},
		}
		var probs = []float64{
			0.3,
			0.4,
			0.3,
			0.3,
		}
		mondrian.DrawR(x0, x1, y0, y1, cmplx, true, cols, probs, "img/mondrian_image.png")
	} else {
		fmt.Println(fmt.Errorf("error: style %s not allowed", style))
		w.WriteHeader(http.StatusInternalServerError)
	}

	htmls := "<html><body><h1>Your Mondrian art : " +
		"</h1><img src='img/mondrian_image.png' width='500px' style='border:1px solid black;'>"

	w.Write([]byte(fmt.Sprintf(htmls)))
}
