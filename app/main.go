package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/user/{name}", getName)
	r.HandleFunc("/magic-number/{num}", magicNumber)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
