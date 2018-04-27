package main

import (
	"fmt"
)

// The (al)most shortest and most easiest way
const (
	dist = 1000
	KB   = dist * iota
	MB   = dist * KB
	GB   = dist * MB
	TB   = dist * GB
	PB   = dist * TB
	EB   = dist * PB
	ZB   = dist * EB
	YB   = dist * ZB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// fmt.Println(ZB)
	// fmt.Println(YB)
}
