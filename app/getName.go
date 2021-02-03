package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func getName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	io.WriteString(w, v["name"])
}
