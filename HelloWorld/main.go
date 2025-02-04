package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Register a handler for the specific path
	// http.ResponseWritter configures http response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! You have requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
