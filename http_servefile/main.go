package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", mazza)
	http.HandleFunc("/mazza_400x400.jpg", mazzapic)
	http.ListenAndServe(":8080", nil)
}
func mazza(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/mazza_400x400.jpg">`)
}

func mazzapic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, `"C:\Users\Drako\golang-web-dev\httpservingalocalfile\mazza_400x400.jpg`)

}
