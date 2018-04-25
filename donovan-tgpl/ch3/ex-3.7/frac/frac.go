package frac

import (
	"image/color"
)

func SupersampledPixel(fr func(z complex128) color.Color, multiplier int, width, height int, pointx, pointy int) color.Color {
	var avgColor color.RGBA
	total := 0 // up to multiplier * multiplier

	bx := pointx * multiplier
	by := pointy * multiplier
	for dxp := 0; dxp < multiplier; dxp++ {
		for dyp := 0; dyp < multiplier; dyp++ {
			col_ := Pixel(fr, width*multiplier, height*multiplier, bx+dxp, by+dyp)
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

func Pixel(fr func(z complex128) color.Color, width, height int, pointx, pointy int) color.Color {
	x := float64(pointx)/float64(width)*(xmax-xmin) + xmin
	y := float64(pointy)/float64(height)*(ymax-ymin) + ymin
	z := complex(x, y)

	return fr(z)
}
