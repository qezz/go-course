package main

import (
	"bytes"
	"fmt"
)

// comma inserts commas in a non-negative decimal integer string.
func RecursiveComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return RecursiveComma(s[:n-3]) + "," + s[n-3:]
}

func IterativeComma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	firstCommaAt := n % 3
	if n < 3 {
		return s
	}
	if firstCommaAt == 0 {
		firstCommaAt = 3
	}
	buf.WriteString(s[:firstCommaAt])

	for i := 1; firstCommaAt+i*3 <= n; i++ {
		buf.WriteByte(',')
		buf.WriteString(s[firstCommaAt+(i-1)*3 : firstCommaAt+i*3])
	}

	return buf.String()
}

func main() {
	fmt.Println(IterativeComma("1123456780"))
}
