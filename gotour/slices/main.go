package main

import (
	"golang.org/x/tour/pic"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	r := [][]uint8{}

	rows := make([]uint8, dy)
	for y := range rows {
		row := make([]uint8, dy)
		for x := range row {
			// row[x] = uint8((x - y*y) / 2)
			row[x] = uint8((x - int(float64(y)*math.Sqrt(float64(y)))) / 2)
		}

		r = append(r, row)
	}

	return r
}

func main() {
	pic.Show(Pic)
}
