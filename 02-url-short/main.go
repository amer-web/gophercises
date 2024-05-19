package main

import (
	"fmt"
	"net/http"
	"os"
	"url-short/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/amer":    "https://godoc.org/github.com/gophercises/urlshort",
		"/mohamed": "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	file, er := os.ReadFile("short.yaml")
	if er != nil {
		fmt.Println("not found", er.Error())
	}

	yamlHandler, err := urlshort.YAMLHandler(file, mapHandler)
	if err != nil {
		fmt.Println(err.Error())
		return
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
