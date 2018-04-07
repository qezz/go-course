package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/qezz/go-course/donovan-tgpl/ch1/ex-1.10/httpcache"
)

// type HttpCache struct {
// 	m sync.Map // map[string]string
// }

func main() {
	cache := httpcache.NewCacheFromFile(".cache")

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
	cache.SaveToFile(".cache")
}

func fetch(url string, cache *httpcache.HttpCache, ch chan<- string) {
	start := time.Now()

	var nbytes int64

	if dom, ispresent := cache.Load(url); ispresent {
		// read from cache
		nbytes2, err := fmt.Fprint(ioutil.Discard, dom.(string))
		if err != nil {
			ch <- fmt.Sprintf("while reading %s from cache: %v", url, err)
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
			ch <- fmt.Sprintf("while reading %s from live: %v", url, err)
			return
		}
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
