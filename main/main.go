package main

import (
	"fmt"
	"net/http"

	"github.com/jonathandly/url_shortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/url_shortener-godoc": "https://godoc.org/github.com/jonathandly/url_shortener",
		"/yaml-godoc":          "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
		- path: /urlshort 
		  url: https://github.com/jonathandly/url_shortener 
		- path: /url_shortener-final
		  url: https://github.com/jonathandly/url_shortener/tree/solution
	`

	yamlHandler, err := url_shortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
