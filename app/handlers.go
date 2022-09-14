package main

import (
	"fmt"
	"image/color"
	"mondrianart/mondrian"
	"net/http"
	"strconv"
)

type PageData struct {
	Complexity    float64
	IsRectandLine bool
	IsRect        bool
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	abouttpl.Execute(w, nil)
}

// imageHandler responds to an incoming request by generating a Mondrian image fo size 500x500
// and according to the chosen style and complexity which is taken from the request form. The
// response is a html block with the image source and dimensions.
func imageHandler(w http.ResponseWriter, r *http.Request) {
	x0 := 0
	x1 := 500
	y0 := 0
	y1 := 500
	complexity := 0.01
	image_location := "img/mondrian_image.png"
	templatedata := PageData{
		Complexity:    complexity,
		IsRectandLine: false,
		IsRect:        true,
	}

	if r.Method == "GET" {
		// render a default image
		mondrian.Draw(x0, x1, y0, y1, complexity, image_location)

	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cmplx, _ := strconv.ParseFloat(r.FormValue("complexity"), 64)
		complexity = cmplx // overwriting to pass later to the template
		style := r.FormValue("style")

		if style == "rectandline" {
			mondrian.Draw(x0, x1, y0, y1, cmplx, image_location)
			templatedata = PageData{
				Complexity:    complexity,
				IsRectandLine: true,
				IsRect:        false,
			}
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
			mondrian.DrawR(x0, x1, y0, y1, cmplx, true, cols, probs, image_location)
			templatedata = PageData{
				Complexity:    complexity,
				IsRectandLine: false,
				IsRect:        true,
			}
		} else {
			fmt.Println(fmt.Errorf("error: style %s not allowed", style))
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	tpl.Execute(w, templatedata)
}
