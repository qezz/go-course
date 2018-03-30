package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot compute Sqrt of a negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 100; i++ {
		prevz := z
		z -= (z*z - x) / (2 * z)

		// for the science
		if math.Abs(z-prevz) < 1e-15 {
			return z, nil
		}
	}

	return z, nil
}

func main() {
	s1, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s1)

	s2, err := Sqrt(-2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s2)
}
