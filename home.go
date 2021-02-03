package main

import (
	"io"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}
