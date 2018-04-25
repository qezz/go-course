package frac

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"strconv"
)

func (frac Frac) SupersampledPixel(fr func(z complex128) color.Color, multiplier int, width, height int, pointx, pointy int) color.Color {
	var avgColor color.RGBA
	total := 0 // up to multiplier * multiplier

	bx := pointx * multiplier
	by := pointy * multiplier
	for dxp := 0; dxp < multiplier; dxp++ {
		for dyp := 0; dyp < multiplier; dyp++ {
			col_ := frac.Pixel(fr, width*multiplier, height*multiplier, bx+dxp, by+dyp)
			col := ColorToRGBA(col_)
			avgColor = color.RGBA{
				uint8((int(avgColor.R)*total + int(col.R)) / (total + 1)),
				uint8((int(avgColor.G)*total + int(col.G)) / (total + 1)),
				uint8((int(avgColor.B)*total + int(col.B)) / (total + 1)),
				uint8((int(avgColor.A)*total + int(col.A)) / (total + 1)),
			}

			total += 1
		}
	}

	return avgColor
}

func ColorToRGBA(col color.Color) color.RGBA {
	r, g, b, a := col.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func (frac Frac) Pixel(fr func(z complex128) color.Color, width, height int, pointx, pointy int) color.Color {
	x := float64(pointx)/float64(width)*(frac.xmax-frac.xmin) + frac.xmin
	y := float64(pointy)/float64(height)*(frac.ymax-frac.ymin) + frac.ymin
	z := complex(x, y)

	return fr(z)
}

type Frac struct {
	height int
	width  int
	ss     int

	xmin float64
	xmax float64
	ymin float64
	ymax float64

	function func(z complex128) color.Color
}

func FracBuilder(vars map[string]string) Frac {
	var height, width, ss int
	var xmin, xmax, ymin, ymax float64
	var f func(z complex128) color.Color

	var err error

	fmt.Println(vars)

	if height, err = strconv.Atoi(vars["height"]); err != nil {
		log.Print(err)
		height = 1024
	}

	if width, err = strconv.Atoi(vars["width"]); err != nil {
		log.Print(err)
		width = 1024
	}

	if ss, err = strconv.Atoi(vars["ss"]); err != nil {
		log.Print(err)
		ss = 1
	}

	if xmin, err = strconv.ParseFloat(vars["xmin"], 64); err != nil {
		log.Print(err)
		xmin = -2.0
	}

	if xmax, err = strconv.ParseFloat(vars["xmax"], 64); err != nil {
		log.Print(err)
		xmax = 2.0
	}

	if ymin, err = strconv.ParseFloat(vars["ymin"], 64); err != nil {
		log.Print(err)
		ymin = -2.0
	}

	if ymax, err = strconv.ParseFloat(vars["ymax"], 64); err != nil {
		log.Print(err)
		ymax = 2.0
	}

	switch vars["f"] {
	case "mandelbrot":
		f = Mandelbrot
	case "newton":
		f = Newton
	default:
		f = Mandelbrot
	}

	return Frac{
		height, width, ss,
		xmin, xmax, ymin, ymax,
		f,
	}
}

func (frac Frac) Write(out io.Writer) {
	width, height := frac.width, frac.height

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			ssp := frac.SupersampledPixel(Mandelbrot, frac.ss, width, height, px, py)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, ssp)
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}
