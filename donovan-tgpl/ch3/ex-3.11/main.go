package main

import (
	"bytes"
	"fmt"
	"strings"
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

	if len(s) == 0 {
		return s
	}

	if s[0] == '+' {
		buf.WriteByte('+')
		s = s[1:]
	} else if s[0] == '-' {
		buf.WriteByte('-')
		s = s[1:]
	}

	n := len(s)

	split := strings.Index(s, ".")

	if split < 0 {
		split = n
	}

	firstCommaAt := split % 3
	if n < 3 {
		buf.WriteString(s)
		return buf.String()
	}
	if firstCommaAt == 0 {
		firstCommaAt = 3
	}
	buf.WriteString(s[:firstCommaAt])

	for i := 1; firstCommaAt+i*3 <= split; i++ {
		buf.WriteByte(',')
		buf.WriteString(s[firstCommaAt+(i-1)*3 : firstCommaAt+i*3])
	}

	if split < n {
		buf.WriteString(s[split:])
	}
	return buf.String()
}

func main() {
	fmt.Println(IterativeComma("1.0123"))
}
