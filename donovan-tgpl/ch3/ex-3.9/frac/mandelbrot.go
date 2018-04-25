package frac

import (
	"image/color"
	"math/cmplx"
)

func Mandelbrot(z complex128) color.Color {
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
