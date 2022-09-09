// package mondrianprocess
package main

import (
	"gonum.org/v1/gonum/stat/distuv"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"
)

var colors = []color.Color{
	color.White,
	// color.Black,
	//color.RGBA{R: 255, G: 0, B: 0, A: 1},
	//color.RGBA{R: 0, G: 255, B: 0, A: 1},
	//color.RGBA{R: 0, G: 0, B: 255, A: 1},
	// red
	color.RGBA{255, 0, 0, 255},
	// yellow
	color.RGBA{255, 255, 0, 255},
	// blue
	color.RGBA{0, 0, 255, 255},
}

type mProcess struct {
	lambda float64
	x0     int
	x1     int
	y0     int
	y1     int
}

func mondrianProcess(mp mProcess, pr *[]image.Rectangle, horizontal bool, complexity float64) *[]image.Rectangle {

	if mp.lambda < 0 || (mp.x1-mp.x0 <= 1 && mp.y1-mp.y0 <= 1) || (mp.x1-mp.x0 == 0 || mp.y1-mp.y0 == 0) {
		// Stop recursive function and return final Rectangles slice
		return pr
	}
	if mp.x1 < mp.x0 {
		log.Fatalln("x1 cannot be smaller than x0")
	}
	if mp.y1 < mp.y0 {
		log.Fatalln("y1 cannot be smaller than y0")
	}
	rand.Seed(time.Now().Unix())

	var mp1 mProcess
	var mp2 mProcess
	//cutRatio := 0.5
	//cut := rand.Float64()
	//fmt.Println(cut)
	//costDist := distuv.Exponential{
	//	Rate: float64(mp.x1 - mp.x0 + mp.y1 - mp.y0),
	//}
	alpha := float64(mp.x1 - mp.x0 + mp.y1 - mp.y0)
	beta := float64(mp.y1-mp.y0) * complexity
	costDist := distuv.Gamma{
		Alpha: alpha,
		Beta:  beta,
	}

	//costDist := distuv.Poisson{
	//	Lambda: float64(mp.x1 - mp.x0 + mp.y1 - mp.y0),
	//}

	// Calculating the new lambda as the old lambda minus the cost of the cut
	lambdaCut := mp.lambda - costDist.Rand()

	// Change the cutting if/else; add new function param horizontal bool, which tells if the previous cut
	// was horizontal or vertical; next cut is supposed to be the opposite
	// if (cut >= cutRatio && (mp.x1-mp.x0 > 1)) || ((mp.x1-mp.x0 > 1) && (mp.y1-mp.y0 <= 1)) {
	if (horizontal && (mp.x1-mp.x0 > 1)) || ((mp.x1-mp.x0 > 1) && (mp.y1-mp.y0 <= 1)) {
		// Horizontal cut
		yCut := mp.y0 + rand.Intn(mp.y1-mp.y0) + 1
		// log.Printf("Making a horizontal cut at y=%d\n", yCut)
		mp1 = mProcess{
			lambda: lambdaCut,
			x0:     mp.x0,
			x1:     mp.x1,
			y0:     mp.y0,
			y1:     yCut,
		}
		mp2 = mProcess{
			lambda: lambdaCut,
			x0:     mp.x0,
			x1:     mp.x1,
			y0:     yCut,
			y1:     mp.y1,
		}
		horizontal = false
	} else if (!horizontal && (mp.y1-mp.y0 > 1)) || ((mp.y1-mp.y0 > 1) && (mp.x1-mp.x0 <= 1)) {
		// Vertical cut
		xCut := mp.x0 + rand.Intn(mp.x1-mp.x0) + 1
		// log.Printf("Making a horizontal cut at x=%d\n", xCut)
		mp1 = mProcess{
			lambda: lambdaCut,
			x0:     mp.x0,
			x1:     xCut,
			y0:     mp.y0,
			y1:     mp.y1,
		}
		mp2 = mProcess{
			lambda: lambdaCut,
			x0:     xCut,
			x1:     mp.x1,
			y0:     mp.y0,
			y1:     mp.y1,
		}
		horizontal = true
	} else {
		return pr
	}

	// Appending the cuts to Rectangles
	*pr = append(
		*pr,
		image.Rectangle{Min: image.Pt(mp1.x0, mp1.y0), Max: image.Pt(mp1.x1, mp1.y1)},
		image.Rectangle{Min: image.Pt(mp2.x0, mp2.y0), Max: image.Pt(mp2.x1, mp2.y1)},
	)

	pr = mondrianProcess(mp1, pr, horizontal, complexity)
	return mondrianProcess(mp2, pr, horizontal, complexity)
}
