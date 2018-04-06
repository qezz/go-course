// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	for _, url := range os.Args[1:] {

		r, _ := regexp.Compile("\\w+://.+") // format: protocol://tail
		if !r.MatchString(url) {
			url = "http://" + url
		}
		fmt.Println("fetching:", url)

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
