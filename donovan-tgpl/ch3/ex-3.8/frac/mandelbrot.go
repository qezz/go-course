package frac

import (
	"image/color"
	"math/big"
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

func MandelbrotBigFloat(z ComplexBigFloat) color.Color {
	const iterations = 200
	const contrast = 15

	v := NewComplexBigFloat()
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		v.Mul(v, v)
		v.Add(v, z)

		vabs := Abs(v)
		// if cmplx.Abs(v) > 2 {
		if vabs.Cmp(big.NewFloat(2.0)) > 0 {
			return color.RGBA{0xcc, 255 - contrast*n, 255 - contrast*n, 0xff}
		}
	}
	return color.Black
}
