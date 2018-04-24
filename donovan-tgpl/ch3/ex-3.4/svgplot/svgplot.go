package svgplot

import (
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"strconv"

	"gopkg.in/go-playground/colors.v1"
)

const angle = math.Pi / 6

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// RGBAShift - may have negative fractional RGBA values
type RGBAShift struct {
	R float32
	G float32
	B float32
	A float32
}

type SVGPlot struct {
	height    int64
	width     int64
	highColor color.RGBA
	lowColor  color.RGBA
	step      RGBAShift // color.RGBA

	cells   int64
	xyrange float64
	xyscale float64
	zscale  float64
	angle   float64
}

func colorsRGBAtocolorRGBA(c *colors.RGBAColor) color.RGBA {
	return color.RGBA{
		c.R,
		c.G,
		c.B,
		uint8(c.A),
	}
}

func SVGPlotBuilder(vars map[string]string) SVGPlot {
	var height, width int64
	var highColor, lowColor color.RGBA

	var err error

	fmt.Println(vars)

	if height, err = strconv.ParseInt(vars["height"], 10, 64); err != nil {
		log.Print(err)
		height = 480
	}

	if width, err = strconv.ParseInt(vars["width"], 10, 64); err != nil {
		log.Print(err)
		width = 800
	}

	if hc_, err := colors.Parse(vars["highColor"]); err != nil {
		log.Print(err)
		highColor = color.RGBA{0xff, 0x0, 0x0, 0xff}
	} else {
		highColor = colorsRGBAtocolorRGBA(hc_.ToRGBA())
	}

	if lc_, err := colors.Parse(vars["lowColor"]); err != nil {
		log.Print(err)
		lowColor = color.RGBA{0x0, 0x0, 0xff, 0xff}
	} else {
		lowColor = colorsRGBAtocolorRGBA(lc_.ToRGBA())
	}

	step := RGBAShift{
		(float32(highColor.R) - float32(lowColor.R)) / 255.0,
		(float32(highColor.G) - float32(lowColor.G)) / 255.0,
		(float32(highColor.B) - float32(lowColor.B)) / 255.0,
		(float32(highColor.A) - float32(lowColor.A)) / 255.0,
	}

	return SVGPlot{
		height:    height,
		width:     width,
		highColor: highColor,
		lowColor:  lowColor,
		step:      step,
	}
}

func (svgp SVGPlot) corner(f func(x, y float64) float64, i, j int64) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := svgp.xyrange * (float64(i)/float64(svgp.cells) - 0.5)
	y := svgp.xyrange * (float64(j)/float64(svgp.cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(svgp.width/2) + (x-y)*cos30*svgp.xyscale
	sy := float64(svgp.height/2) + (x+y)*sin30*svgp.xyscale - z*svgp.zscale
	return sx, sy, z * svgp.zscale
}

func (svgp SVGPlot) Write(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", svgp.width, svgp.height)

	svgp.cells = 50
	svgp.xyrange = 10.0
	svgp.xyscale = float64(svgp.width/2) / 10.0
	svgp.zscale = float64(svgp.height) * 0.01
	svgp.angle = math.Pi / 6.0

	// choose a fucntion to vizualize
	f := saddle

	var i, j int64
	for i = 0; i < svgp.cells; i++ {
		for j = 0; j < svgp.cells; j++ {
			var h float64

			ax, ay, az := svgp.corner(f, i+1, j)
			if IsInvalid(ax, ay) {
				continue
			}
			h += az

			bx, by, bz := svgp.corner(f, i, j)
			if IsInvalid(bx, by) {
				continue
			}
			h += bz

			cx, cy, cz := svgp.corner(f, i, j+1)
			if IsInvalid(cx, cy) {
				continue
			}
			h += cz

			dx, dy, dz := svgp.corner(f, i+1, j+1)
			if IsInvalid(dx, dy) {
				continue
			}
			h += dz

			he := (h/4.0 + float64(svgp.height)) / float64(svgp.height*2) * 255.0 // from -1000 to +1000

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'",
				ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(out, " style=\"fill:rgb(%v,%v,%v)\" />\n",
				svgp.lowColor.R+uint8(he*float64(svgp.step.R)),
				svgp.lowColor.G+uint8(he*float64(svgp.step.G)),
				svgp.lowColor.B+uint8(he*float64(svgp.step.B)))
		}
	}

	fmt.Fprintln(out, "</svg>")
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

func IsInvalid(a, b float64) bool {
	if math.IsInf(a, 0) || math.IsNaN(a) || math.IsInf(a, 0) || math.IsNaN(a) {
		return true
	}

	return false
}
