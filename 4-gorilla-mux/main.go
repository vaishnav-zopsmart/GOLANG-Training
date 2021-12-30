package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "User %s\n", vars["id"])
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func name(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Vaishnav!")
}

func api(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	fmt.Fprintf(w, "Name is : %v", name)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello/{id}", hello)
	r.HandleFunc("/world", world)
	r.HandleFunc("/name/{name}", name).Methods(http.MethodGet)
	r.HandleFunc("/api", api).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
