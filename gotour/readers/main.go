package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	l := len(b)
	for i := 0; i < l; i++ {
		b[i] = 'A'
	}

	return l, nil
}

func main() {
	reader.Validate(MyReader{})
}
