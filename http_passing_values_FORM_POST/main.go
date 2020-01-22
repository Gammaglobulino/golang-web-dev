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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="get">
								<input type="text" name="q">
								<input type="submit">
							</form>
							<br>`+v)
}
