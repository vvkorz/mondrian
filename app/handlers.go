package main

import (
	"fmt"
	"image/color"
	"mondrianart/mondrian"
	"net/http"
	"strconv"
)

type PageData struct {
	Complexity float64
	Lines      bool
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
		Complexity: complexity,
		Lines:      false,
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
		var cols = []color.Color{
			color.White,
			color.RGBA{34, 80, 149, 255}, // blue
			color.RGBA{250, 201, 1, 255}, // yellow
			color.RGBA{221, 1, 0, 255},   // red
			color.RGBA{30, 38, 33, 255},  // black
		}
		var probs = []float64{
			0.5,
			0.13,
			0.13,
			0.13,
			0.11,
		}
		if style == "on" {
			// lines
			mondrian.DrawR(x0, x1, y0, y1, cmplx, false, cols, probs, image_location, 2)
			templatedata = PageData{
				Complexity: complexity,
				Lines:      true,
			}
		} else if style == "" {
			// no lines
			mondrian.DrawR(x0, x1, y0, y1, cmplx, true, cols, probs, image_location, 2)
			templatedata = PageData{
				Complexity: complexity,
				Lines:      false,
			}
		} else {
			fmt.Println(fmt.Errorf("error: style %s not allowed", style))
			w.WriteHeader(http.StatusInternalServerError)
		}

	}

	tpl.Execute(w, templatedata)
}
