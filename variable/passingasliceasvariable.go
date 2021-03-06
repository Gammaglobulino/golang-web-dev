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
	//persons := []string{"Andrea", "luca", "Stefano"}
	/*persons1 := map[string]string{
		"Andrea":"Senior Engineer",
		"Stefano":"Carabiniere",
	}
	*/
	type person struct {
		Name     string
		Lastname string
	}
	ps := []person{
		{"Andrea", "Mazzanti"},
		{"Stefano", "Mazzanti"},
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tplnew.ExecuteTemplate(nf, "tpl_passingacustomtype.gohtml", ps)
	if err != nil {
		log.Fatalln(err)
	}
}
