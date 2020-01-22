package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(`.`)))
	http.HandleFunc("/mazza", mazza)
	http.ListenAndServe(":8080", nil)
}
func mazza(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/mazza_400x400.jpg">`)
}
