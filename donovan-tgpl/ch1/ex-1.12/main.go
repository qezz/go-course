// Server2 is a minimal "echo" and counter server.
package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/qezz/go-course/donovan-tgpl/ch1/ex-1.12/gifs"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	config := make(map[string]string)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		config[k] = v[0]
	}

	lg := gifs.LissajousBuilder(w, config)
	lg.Write(w)
}
