package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/smelton01/go-basics/urlshort"
)

func main() {
	// var yamlFile = flag.String("-yaml", "paths.yml", "Path to yaml file with url shortcuts")
	var filePath = flag.String("-file", "paths.json", "Path to file with url shortcuts")

	flag.Parse()

	mux := defaultMux()

	var handler http.Handler
	file, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal("File error: ", err)
	}

	if ext := filepath.Ext(*filePath); ext == ".yaml" {
		handler, err = urlshort.YAMLHandler(file, mux)

		if err != nil {
			log.Fatal("Handler Error", err)
		}
	} else if ext == ".json" {
		handler, err = urlshort.JSONHandler(file, mux)
		if err != nil {
			log.Fatal("Handler Error", err)
		}
	}else {
		// TO DO: make db default
		handler = mux
	}
	
	// // Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	// mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
