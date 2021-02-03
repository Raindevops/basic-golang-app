package main

import (
	"io"
	"math"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
