package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
	fmt.Fprintf(w, "World!")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	http.ListenAndServe(":8000", nil)
}
