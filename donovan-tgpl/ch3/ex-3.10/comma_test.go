package main

import "testing"

var commaTests []struct {
	left  string
	right string
}

func init() {
	commaTests = []struct {
		left  string
		right string
	}{
		{"", ""},
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"1234", "1,234"},
		{"12345", "12,345"},
		{"123456", "123,456"},
		{"123456789", "123,456,789"},
	}
}

func TestRecursiveComma(t *testing.T) {
	for _, row := range commaTests {
		if RecursiveComma(row.left) != row.right {
			t.Fail()
		}
	}
}

func TestIterativeComma(t *testing.T) {
	for _, row := range commaTests {
		if IterativeComma(row.left) != row.right {
			t.Fail()
		}
	}
}
