package main

import (
	"fmt"
	"net/http"
)

type andrea int

func (m andrea) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Any code you want here")
}
func main() {
	var d andrea
	http.ListenAndServe(":8080", d)
}
