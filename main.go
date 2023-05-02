package main

import (
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About Page"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merhaba Mars!"))
	})

	http.HandleFunc("/about", aboutHandler)

	http.ListenAndServe(":8080", nil)
}
