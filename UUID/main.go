package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", base)
	http.ListenAndServe(":8080", nil) //using GO default handler

}

func base(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
