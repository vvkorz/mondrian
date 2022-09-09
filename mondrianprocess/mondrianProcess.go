// package mondrianprocess
package main

import (
	"gonum.org/v1/gonum/stat/distuv"
	"image"
	"log"
	"math/rand"
	"time"
)

// mProcess defines the parameters of a Mondrian Process: the partitioning cost lambda,
// the horizontal points x0 and x1 and the vertical points y0 and y1.
type mProcess struct {
	lambda float64
	x0     int
	x1     int
	y0     int
	y1     int
}

// mondrianProcess takes an initial mProcess mp and a pointer pr to the initial image canvas
// and recursively creates rectangle partitions under given complexity and
// horizontal/vertical direction.
func mondrianProcess(
	mp mProcess,
	pr *[]image.Rectangle,
	horizontal bool,
	complexity float64,
) *[]image.Rectangle {

	if mp.lambda < 0 ||
		(mp.x1-mp.x0 <= 1 && mp.y1-mp.y0 <= 1) ||
		(mp.x1-mp.x0 == 0 || mp.y1-mp.y0 == 0) {
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
	alpha := float64(mp.x1 - mp.x0 + mp.y1 - mp.y0)
	beta := float64(mp.y1-mp.y0) * complexity
	costDist := distuv.Gamma{
		Alpha: alpha,
		Beta:  beta,
	}
	// Calculating the new lambda as the old lambda minus the cost of the cut
	lambdaCut := mp.lambda - costDist.Rand()

	if (horizontal && (mp.x1-mp.x0 > 1)) ||
		((mp.x1-mp.x0 > 1) && (mp.y1-mp.y0 <= 1)) {
		// Horizontal cut
		yCut := mp.y0 + rand.Intn(mp.y1-mp.y0) + 1
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
	} else if (!horizontal && (mp.y1-mp.y0 > 1)) ||
		((mp.y1-mp.y0 > 1) && (mp.x1-mp.x0 <= 1)) {
		// Vertical cut
		xCut := mp.x0 + rand.Intn(mp.x1-mp.x0) + 1
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
	// Recursively starting the function for the two new partitions
	pr = mondrianProcess(mp1, pr, horizontal, complexity)
	return mondrianProcess(mp2, pr, horizontal, complexity)
}
