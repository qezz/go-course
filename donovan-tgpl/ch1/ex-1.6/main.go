// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{
	// color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.Black,
	// color.RGBA{0x0, 0xff, 0x0, 0xff},
	// color.RGBA{0xa, 0xcc, 0xf, 0xff},
	color.RGBA{0xff, 0x0, 0x0, 0xff},  // red
	color.RGBA{0xff, 0x7f, 0x0, 0xff}, // orange
	color.RGBA{0xff, 0xff, 0x0, 0xff}, // yellow
	color.RGBA{0x7f, 0xff, 0x0, 0xff}, // spring green
	color.RGBA{0x0, 0xff, 0x0, 0xff},  // green
	color.RGBA{0x0, 0xff, 0x7f, 0xff}, // turquoise
	color.RGBA{0x0, 0xff, 0xff, 0xff}, // cyan
	color.RGBA{0x0, 0x7f, 0xff, 0xff}, // ocean
	color.RGBA{0x0, 0x0, 0xff, 0xff},  // blue
	color.RGBA{0x7f, 0x0, 0xff, 0xff}, // violet
	color.RGBA{0xff, 0x0, 0xff, 0xff}, // magenta
	color.RGBA{0xff, 0x0, 0x7f, 0xff}, // raspberry
}

const (
	blackIndex = 0
	greenIndex = 1
)

func main() {
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200
		nframes = 64
		delay   = 5
		// image canvas covers [-size..+size]
		// number of animation frames
		// delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%(len(palette)-1)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
