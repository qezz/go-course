package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/qezz/go-course/donovan-tgpl/ch2/ex-2.3/popcount"
)

func XorArrays(a, b [32]uint8) []uint8 {
	if len(a) != len(b) {
		return nil
	}

	res := make([]uint8, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}

	return res
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	x := XorArrays(c1, c2)
	acc := 0
	for i := 0; i < len(x); i++ {
		acc += popcount.PopCount2(uint64(x[i]))
	}
	fmt.Println(acc)
}
