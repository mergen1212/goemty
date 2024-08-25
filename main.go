package main

import (
	"io"
	"net/http"
)

const s = "Home page!"

func r(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, s)
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", r)
	http.ListenAndServe(":8080", mux)
}
