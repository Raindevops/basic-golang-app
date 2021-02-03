package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func greet(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func getName(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Get name func")
}

func magicNumber(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Magic number function")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", greet)
	r.HandleFunc("/user/{name}", getName)
	r.HandleFunc("/magic-number/{num}", magicNumber)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
