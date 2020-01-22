package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {

	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", handleForm)
	http.ListenAndServe(":8080", nil)
}

func handleForm(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribed") == "on"

	err := tpl.ExecuteTemplate(w, "index.gohtml", person{
		FirstName:  f,
		LastName:   l,
		Subscribed: s,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatal(err)
	}
}
