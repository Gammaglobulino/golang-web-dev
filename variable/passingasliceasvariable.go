package main

import (
	"log"
	"os"
	"text/template"
)

var tplnew *template.Template

func init() {
	tplnew = template.Must(template.ParseGlob("variable/*"))
}
func main() {
	persons := []string{"Andrea", "luca", "Stefano"}
	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tplnew.ExecuteTemplate(nf, "tpl_with_rangeindex.gohtml", persons)
	if err != nil {
		log.Fatalln(err)
	}
}
