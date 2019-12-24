package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("variable/*"))
}
func main() {

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tpl.ExecuteTemplate(nf, "tpl.gohtml", "release self focus")
	if err != nil {
		log.Fatalln(err)
	}
}
