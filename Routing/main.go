package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// mux provides more powerful handle function compared to http.HandleFunc.
	//It can match path parameters and create multiple router sets
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You have requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":8080", r)
}
