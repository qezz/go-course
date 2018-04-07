package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync" // sync map
	"time"

	"bytes"
)

// type HttpCache struct {
// 	m sync.Map // map[string]string
// }

type HttpCache struct {
	sync.Map
}

func (f *HttpCache) UnmarshalJSON(data []byte) error {
	var tmpMap map[string]interface{}
	if err := json.Unmarshal(data, &tmpMap); err != nil {
		return err
	}
	for key, value := range tmpMap {
		f.Store(key, value)
	}
	return nil
}

func (c HttpCache) MarshalJSON() ([]byte, error) {
	tmpMap := make(map[string]string)
	c.Range(func(k, v interface{}) bool {
		tmpMap[k.(string)] = v.(string)
		return true
	})
	return json.Marshal(tmpMap)
}

func main() {
	// load cache if present
	var cache HttpCache

	dat, err := ioutil.ReadFile(".cache")
	if err == nil {
		err2 := json.Unmarshal(dat, &cache)
		if err2 != nil {
			fmt.Fprintln(os.Stderr, "Can't parse json")
		}
	}

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, &cache, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	// save to file

	b, err := json.Marshal(cache)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	f, err := os.Create(".cache")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer f.Close()

	f.Write(b)

}

func fetch(url string, cache *HttpCache, ch chan<- string) {
	start := time.Now()

	var nbytes int64

	if dom, ispresent := cache.Load(url); ispresent {
		// read from cache
		nbytes2, err := fmt.Fprintln(ioutil.Discard, dom.(string))
		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
		}
		nbytes = int64(nbytes2)
	} else {
		resp, err := http.Get(url)
		defer resp.Body.Close() // don't leak resources
		if err != nil {
			ch <- fmt.Sprint(err) // send to channel ch
			return
		}

		buf := new(bytes.Buffer)
		nbytes, err = io.Copy(buf, resp.Body)
		cache.Store(url, buf.String())

		if err != nil {
			ch <- fmt.Sprintf("while reading %s: %v", url, err)
			return
		}
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
