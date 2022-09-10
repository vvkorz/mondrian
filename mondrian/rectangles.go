package mondrian

import (
	"gonum.org/v1/gonum/stat/distuv"
	"image"
	"log"
	"math/rand"
	"time"
)

// MProcess defines the parameters of a Mondrian Process: the partitioning cost lambda,
// the horizontal points x0 and x1 and the vertical points y0 and y1.
type MProcess struct {
	Lambda float64
	X0     int
	X1     int
	Y0     int
	Y1     int
}

// MondrianProcess takes an initial mProcess mp and a pointer pr to the initial image canvas
// and recursively creates rectangle partitions under given complexity and
// horizontal/vertical direction.
func MondrianProcess(
	mp MProcess,
	pr *[]image.Rectangle,
	horizontal bool,
	complexity float64,
) *[]image.Rectangle {

	if mp.Lambda < 0 ||
		(mp.X1-mp.X0 <= 1 && mp.Y1-mp.Y0 <= 1) ||
		(mp.X1-mp.X0 == 0 || mp.Y1-mp.Y0 == 0) {
		// Stop recursive function and return final Rectangles slice
		return pr
	}
	if mp.X1 < mp.X0 {
		log.Fatalln("x1 cannot be smaller than x0")
	}
	if mp.Y1 < mp.Y0 {
		log.Fatalln("y1 cannot be smaller than y0")
	}
	rand.Seed(time.Now().Unix())

	var mp1 MProcess
	var mp2 MProcess
	alpha := float64(mp.X1 - mp.X0 + mp.Y1 - mp.Y0)
	beta := float64(mp.Y1-mp.Y0) * complexity
	costDist := distuv.Gamma{
		Alpha: alpha,
		Beta:  beta,
	}
	// Calculating the new lambda as the old lambda minus the cost of the cut
	lambdaCut := mp.Lambda - costDist.Rand()

	if (horizontal && (mp.X1-mp.X0 > 1)) ||
		((mp.X1-mp.X0 > 1) && (mp.Y1-mp.Y0 <= 1)) {
		// Horizontal cut
		yCut := mp.Y0 + rand.Intn(mp.Y1-mp.Y0) + 1
		mp1 = MProcess{
			Lambda: lambdaCut,
			X0:     mp.X0,
			X1:     mp.X1,
			Y0:     mp.Y0,
			Y1:     yCut,
		}
		mp2 = MProcess{
			Lambda: lambdaCut,
			X0:     mp.X0,
			X1:     mp.X1,
			Y0:     yCut,
			Y1:     mp.Y1,
		}
		horizontal = false
	} else if (!horizontal && (mp.Y1-mp.Y0 > 1)) ||
		((mp.Y1-mp.Y0 > 1) && (mp.X1-mp.X0 <= 1)) {
		// Vertical cut
		xCut := mp.X0 + rand.Intn(mp.X1-mp.X0) + 1
		mp1 = MProcess{
			Lambda: lambdaCut,
			X0:     mp.X0,
			X1:     xCut,
			Y0:     mp.Y0,
			Y1:     mp.Y1,
		}
		mp2 = MProcess{
			Lambda: lambdaCut,
			X0:     xCut,
			X1:     mp.X1,
			Y0:     mp.Y0,
			Y1:     mp.Y1,
		}
		horizontal = true
	} else {
		return pr
	}

	// Appending the cuts to Rectangles
	*pr = append(
		*pr,
		image.Rectangle{Min: image.Pt(mp1.X0, mp1.Y0), Max: image.Pt(mp1.X1, mp1.Y1)},
		image.Rectangle{Min: image.Pt(mp2.X0, mp2.Y0), Max: image.Pt(mp2.X1, mp2.Y1)},
	)
	// Recursively starting the function for the two new partitions
	pr = MondrianProcess(mp1, pr, horizontal, complexity)
	return MondrianProcess(mp2, pr, horizontal, complexity)
}
