package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/qezz/go-course/donovan-tgpl/ch2/ex-2.1/tempconv"
)

func doAllConversions(value string) {
	t, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			doAllConversions(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			doAllConversions(arg)
		}

	}
}
