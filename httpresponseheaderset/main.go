package main

import (
	"fmt"
	"net/http"
)

type andrea int

//HTTP handler
func (a andrea) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Andrea-Mazzanti", "yes")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> Ciao bello come stai?</h1>")
}

func main() {
	var a andrea
	http.ListenAndServe(":8080", a)
}
