// Exercise 1.1: Modify the echo program to also print os.Args[0], the
// name of the command that invoked it.

// package ch1ex1dot1

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		fmt.Printf("[%v] %v\n", i, a)
	}
}
