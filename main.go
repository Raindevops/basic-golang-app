package main

import (
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func getName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	io.WriteString(w, v["name"])
}

func magicNumber(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	n, err := strconv.Atoi(v["num"])
	if err != nil {
		io.WriteString(w, "Please enter a valide number")
	} else {
		mn := math.Pow(float64(n), 3)
		io.WriteString(w, strconv.Itoa(int(mn)))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/user/{name}", getName)
	r.HandleFunc("/magic-number/{num}", magicNumber)

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
