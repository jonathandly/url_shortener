package main

import (
	"fmt"
	"net/http"

	"github.com/jonathandly/urlshortener"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshortener-godoc": "https://godoc.org/github.com/jonathandly/urlshortener",
		"/yaml-godoc":         "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
		- path: /urlshortener 
		  url: https://github.com/jonathandly/urlshortener 
		- path: /url_shortener-final
		  url: https://github.com/jonathandly/urlshortener/tree/tut
	`

	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
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
