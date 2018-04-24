// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	// "fmt"
	"log"
	"math"
	"net/http"

	"github.com/qezz/go-course/donovan-tgpl/ch3/ex-3.4/svgplot"
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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	config := make(map[string]string)

	w.Header().Set("Content-Type", "image/svg+xml")

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		config[k] = v[0]
	}

	svg := svgplot.SVGPlotBuilder(config)
	svg.Write(w)
}
