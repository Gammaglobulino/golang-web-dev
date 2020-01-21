package main

import (
	"io"
	"net/http"
)

type andrea int
type dog int
type cat int

//HTTP handler rough style
func (a andrea) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(w, "whof whof whof")
	case "/cat":
		io.WriteString(w, "miau miau miau")

	}
}

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "whof whof whof")

}

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "miau miau miau")

}

func main() {
	//var a andrea

	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)
	http.ListenAndServe(":8080", nil)
}
