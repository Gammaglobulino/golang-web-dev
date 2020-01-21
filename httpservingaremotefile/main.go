package main

import (
	"io"
	"net/http"
)

func dogs(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	io.WriteString(w, `<img alt="Collage of Nine Dogs.jpg" src="//upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Collage_of_Nine_Dogs.jpg/260px-Collage_of_Nine_Dogs.jpg" decoding="async" width="260" height="228" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Collage_of_Nine_Dogs.jpg/390px-Collage_of_Nine_Dogs.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/d/d9/Collage_of_Nine_Dogs.jpg/520px-Collage_of_Nine_Dogs.jpg 2x" data-file-width="1665" data-file-height="1463">`)
}

func main() {
	http.HandleFunc("/", dogs)
	http.ListenAndServe(":8080", nil)
}
