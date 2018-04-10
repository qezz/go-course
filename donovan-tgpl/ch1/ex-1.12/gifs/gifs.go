// Lissajous generates GIF animations of random Lissajous figures.
package gifs

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{
	// color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	// color.Black,
	color.RGBA{0xfd, 0xfd, 0xfd, 0xff},
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

type LissajousGif struct {
	cycles  int
	res     float64
	size    int
	nframes int
	delay   int
}

func LissajousBuilder(out io.Writer, vars map[string]string) LissajousGif {
	var cycles, size, nframes, delay int
	var res float64
	var err error

	if cycles, err = strconv.Atoi(vars["cycles"]); err != nil {
		log.Print(err)
		cycles = 5
	}

	if res, err = strconv.ParseFloat(vars["res"], 64); err != nil {
		log.Print(err)
		res = 0.001
	}

	if size, err = strconv.Atoi(vars["size"]); err != nil {
		log.Print(err)
		size = 200
	}

	if nframes, err = strconv.Atoi(vars["nframes"]); err != nil {
		log.Print(err)
		nframes = 200
	}

	if delay, err = strconv.Atoi(vars["delay"]); err != nil {
		log.Print(err)
		delay = 20
	}

	fmt.Println("size:", size)

	return LissajousGif{
		cycles:  cycles,
		res:     res,
		size:    size,
		nframes: nframes,
		delay:   delay,
	}
}

func (lg LissajousGif) Write(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: lg.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < lg.nframes; i++ {
		rect := image.Rect(0, 0, 2*lg.size+1, 2*lg.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(lg.cycles)*2*math.Pi; t += lg.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(lg.size+int(x*float64(lg.size)+0.5),
				lg.size+int(y*float64(lg.size)+0.5),
				uint8(i%(len(palette)-1)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, lg.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
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
