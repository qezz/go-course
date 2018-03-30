package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	r := map[string]int{}
	words := strings.Fields(s)
	for _, word := range words {
		r[word] += 1
	}
	return r
}

func main() {
	wc.Test(WordCount)
}
