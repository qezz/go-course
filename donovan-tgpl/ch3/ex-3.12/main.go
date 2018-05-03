package main

import (
	"fmt"
	"sync"
)

func WordFreqAnalizer(phrase string) map[rune]int {
	dict := make(map[rune]int)

	for _, letter := range phrase {
		// skip spaces
		if letter == rune(' ') {
			continue
		}
		dict[letter] += 1
	}

	return dict
}

func AreMapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}

	return true
}

func AsyncFreqsToHashMap(phrase string, m *map[rune]int, wg *sync.WaitGroup) {
	*m = WordFreqAnalizer(phrase)
	wg.Done()
}

func AreAnagrams(a, b string) bool {
	aDict := make(map[rune]int)
	bDict := make(map[rune]int)

	var wg sync.WaitGroup
	wg.Add(2)
	go AsyncFreqsToHashMap(a, &aDict, &wg)
	go AsyncFreqsToHashMap(b, &bDict, &wg)
	wg.Wait()

	// Someone said that reflect.DeepEqual() is about 80 times
	// slower than custom comparition function.
	return AreMapsEqual(aDict, bDict)
}

func main() {
	fmt.Println(":", AreAnagrams("b a", "ab"))
}
