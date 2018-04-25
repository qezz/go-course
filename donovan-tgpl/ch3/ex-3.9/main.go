// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"log"
	"net/http"

	"github.com/qezz/go-course/donovan-tgpl/ch3/ex-3.9/frac"
)

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

	frac := frac.FracBuilder(config)
	frac.Write(w)
}
