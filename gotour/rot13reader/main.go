package main

import (
	// "fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func inRange(char, lower, upper byte) bool {
	return char >= lower && char <= upper
}

func trasformRotChar(char byte, base, rot uint8, lower byte) byte {
	return (char-lower+rot)%base + lower
}

func (self rot13Reader) Read(output []byte) (int, error) {
	haveread, err := self.r.Read(output)

	for i := 0; i < haveread; i++ {
		// // before:
		// if output[i] >= 'A' && output[i] <= 'Z' {
		// 	output[i] = (output[i]-'A'+13)%26 + 'A'
		// } else if output[i] >= 'a' && output[i] <= 'z' {
		// 	output[i] = (output[i]-'a'+13)%26 + 'a'
		// }

		// after:
		if inRange(output[i], 'A', 'Z') {
			output[i] = trasformRotChar(output[i], 26, 13, 'A')
		} else if inRange(output[i], 'a', 'z') {
			output[i] = trasformRotChar(output[i], 26, 13, 'a')
		}
	}

	if err == io.EOF {
		return haveread, err
	}

	return haveread, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
