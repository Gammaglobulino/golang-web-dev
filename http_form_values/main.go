package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", mazza)
	http.ListenAndServe(":8080", nil)
}
func mazza(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	io.WriteString(w, `Do my search: `+v)
}
