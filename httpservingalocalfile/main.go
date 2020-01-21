package main

import (
	"io"
	"net/http"
	"os"
)

func mazza(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/mazza_400x400.jpg">`)
}

func mazzapic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open(`C:\Users\Drako\golang-web-dev\httpservingalocalfile\mazza_400x400.jpg`)
	if err != nil {
		http.Error(w, "file not found", 404)
	}
	defer f.Close()
	io.Copy(w, f)

}
func main() {
	http.HandleFunc("/", mazza)
	http.HandleFunc("/mazza_400x400.jpg", mazzapic)
	http.ListenAndServe(":8080", nil)
}
