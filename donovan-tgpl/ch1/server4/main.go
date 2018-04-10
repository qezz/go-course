// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/qezz/go-course/donovan-tgpl/ch1/ex-1.6"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	lissajous(w)
}
