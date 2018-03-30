package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f1, f2 = f2, f1 + f2
		
		return f1    // returns 1 1 2 3 5 ...
		// return f2 // returns 1 2 3 5 8 ...
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

