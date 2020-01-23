package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(`http_riderects/templates/*`))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", base)
	http.HandleFunc("/here", here)
	http.HandleFunc("/there", there)

	http.ListenAndServe(":8080", nil) //using GO default handler

}

func base(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at base", req.Method, "\n\n")
}
func there(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at there", req.Method, "\n\n")
	w.Header().Set("Location", "/here")
	w.WriteHeader(http.StatusSeeOther) //303
}

func here(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at here", req.Method, "\n\n")
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
