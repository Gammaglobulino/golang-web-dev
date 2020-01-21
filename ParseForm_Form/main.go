package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type andrea int

//HTTP handler
func (a andrea) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	if err := req.ParseForm(); err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("ParseForm_Form/index.gohtml"))
}
func main() {
	var a andrea
	http.ListenAndServe(":8080", a)
}
