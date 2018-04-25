package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

func SupersampledPixel(multiplier int, width, height int, pointx, pointy int) color.Color {
	var avgColor color.RGBA
	total := 0 // up to multiplier * multiplier

	bx := pointx * multiplier
	by := pointy * multiplier
	for dxp := 0; dxp < multiplier; dxp++ {
		for dyp := 0; dyp < multiplier; dyp++ {
			col_ := Pixel(width*multiplier, height*multiplier, bx+dxp, by+dyp)
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

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
)

func ColorToRGBA(col color.Color) color.RGBA {
	r, g, b, a := col.RGBA()
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func Pixel(width, height int, pointx, pointy int) color.Color {
	x := float64(pointx)/float64(width)*(xmax-xmin) + xmin
	y := float64(pointy)/float64(height)*(ymax-ymin) + ymin
	z := complex(x, y)

	return newton(z)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{0xcc, 255 - contrast*n, 255 - contrast*n, 0xff}
		}
	}
	return color.Black
}

// Stolen from gopl, accidentally
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
