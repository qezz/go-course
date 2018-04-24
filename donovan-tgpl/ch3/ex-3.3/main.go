// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

// Saddle setup
const (
	width, height = 800, 480         // canvas size in pixels
	cells         = 50               // number of grid cells
	xyrange       = 10.0             // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / 10.0 // xyrange // pixels per x or y unit
	zscale        = height * 0.01    // pixels per z unit
	angle         = math.Pi / 6      // angle of x, y axes (=30°)
)

// // Setup for egg box:
// const (
// 	width, height = 800, 480         // canvas size in pixels
// 	cells         = 100              // number of grid cells
// 	xyrange       = 25.0             // axis ranges (-xyrange..+xyrange)
// 	xyscale       = width / 2 / 20.0 // xyrange // pixels per x or y unit
// 	zscale        = height * 0.012   // pixels per z unit
// 	angle         = math.Pi / 6      // angle of x, y axes (=30°)
// )

// // Setup for moguls
// const (
// 	width, height = 800, 480         // canvas size in pixels
// 	cells         = 150              // number of grid cells
// 	xyrange       = 15.0             // axis ranges (-xyrange..+xyrange)
// 	xyscale       = width / 2 / 20.0 // xyrange // pixels per x or y unit
// 	zscale        = height * 0.012   // pixels per z unit
// 	angle         = math.Pi / 6      // angle of x, y axes (=30°)
// )

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	// choose a fucntion to vizualize
	f := saddle

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var h float64

			ax, ay, az := corner(f, i+1, j)
			if IsInvalid(ax, ay) {
				continue
			}
			h += az

			bx, by, bz := corner(f, i, j)
			if IsInvalid(bx, by) {
				continue
			}
			h += bz

			cx, cy, cz := corner(f, i, j+1)
			if IsInvalid(cx, cy) {
				continue
			}
			h += cz

			dx, dy, dz := corner(f, i+1, j+1)
			if IsInvalid(dx, dy) {
				continue
			}
			h += dz

			he := int(h/4.0 + 125.0) // hack

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'", // />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Printf(" style=\"fill:rgb(%v,0,%v)\" />\n",
				he, 255-he)
		}
	}
	fmt.Println("</svg>")
}

func IsInvalid(a, b float64) bool {
	if math.IsInf(a, 0) || math.IsNaN(a) || math.IsInf(a, 0) || math.IsNaN(a) {
		return true
	}

	return false
}

func corner(f func(x, y float64) float64, i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z * zscale
}

func basic(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return 7 * (math.Sin(x)*math.Sin(x) + math.Sin(y)*math.Sin(y))
}

// Free interpretation of word 'moguls', looks like splitted eggbox
func moguls(x, y float64) float64 {
	return -math.Abs(7*(math.Sin(x)*math.Sin(x)+math.Sin(y)*math.Sin(y)) - 7)
}

func saddle(x, y float64) float64 {
	return x*x - y*y
}
