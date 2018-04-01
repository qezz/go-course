package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (i Image) At(x, y int) color.Color {
	v := uint8((x*y)/(x+y+1) - 1)
	return color.RGBA{uint8(v * uint8(y) / (uint8(x)%255 + 1)), v - uint8(y), v - uint8(x), 211}
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
