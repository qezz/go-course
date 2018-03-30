package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 100; i++ {
		prevz := z
		z -= (z*z - x) / (2 * z)

		// just for fun
		if math.Abs(z-prevz) < 1e-15 {
			return z
		}
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
