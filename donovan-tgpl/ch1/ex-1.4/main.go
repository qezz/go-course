// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

// Do not run with `go run main.go main.go`, because `go run` can't
// distinguish between source files and arguments

// Modify dup2 to print the names of all files in which each
// duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

type WordEntry struct {
	//	name string
	count   int
	sources map[string]int
}

func NewWordEntry() WordEntry {
	return WordEntry{}
}

func (we *WordEntry) AddSource(source string) {
	we.count++
	if we.sources == nil {
		we.sources = make(map[string]int)
	}

	we.sources[source]++
}

func (we WordEntry) SourcesAsString() string {
	s := "[ "
	for k := range we.sources {
		s += k
		s += " "
	}

	s += "]"

	return s
}

func main() {
	counts := make(map[string]WordEntry)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			fp, err := filepath.Abs(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			} else {
				fmt.Printf("read %v\n", fp)
			}

			countLines(f, counts)
			f.Close()
		}
	}
	for line, we := range counts {
		if we.count > 1 {
			fmt.Printf("%d\t%s\t%v\n", we.count, line, we.SourcesAsString())
		}
	}
}

func countLines(f *os.File, counts map[string]WordEntry) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		we, _ := counts[input.Text()]
		we.AddSource(f.Name())

		counts[input.Text()] = we
	}
	// NOTE: ignoring potential errors from input.Err()
}
