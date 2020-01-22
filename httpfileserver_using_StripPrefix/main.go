package main

import (
	"io"
	"net/http"
)

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir(`C:\Users\Drako\golang-web-dev\httpfileserver_using_StripPrefix\assets`))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", mazza)
	http.ListenAndServe(":8080", nil)
}
func mazza(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/mazza_400x400.jpg">`)
}
