package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", base)
	http.HandleFunc("/read", read)

	http.ListenAndServe(":8080", nil) //using GO default handler

}

func base(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "Andrea",
		Value: "Andrea cookie",
	})
	fmt.Fprintf(w, "Cookie-written check your browser")
}
func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("Andrea")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "your cookie:", c)

}
