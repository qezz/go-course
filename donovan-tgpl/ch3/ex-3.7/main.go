// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/png"
	"os"

	"github.com/qezz/go-course/donovan-tgpl/ch3/ex-3.7/frac"
)

func main() {
	const (
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			ssp := frac.SupersampledPixel(frac.Mandelbrot, 10, width, height, px, py)

			// Image point (px, py) represents complex value z.
			img.Set(px, py, ssp)
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}
