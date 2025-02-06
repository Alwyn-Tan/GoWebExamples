package main

import (
	"fmt"
	"log"
	"net/http"
)

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Requested path: %s", r.URL.Path)
		next(w, r)
	}
}

func AlwynHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Alwyn")
}

func IreneHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Irene")
}

func main() {
	http.HandleFunc("/Alwyn", logging(IreneHandler))
	http.HandleFunc("/Irene", logging(AlwynHandler))
	http.ListenAndServe(":8080", nil)
}
