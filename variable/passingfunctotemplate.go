package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tplnew1 *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
}

func init() {
	tplnew1 = template.Must(template.New("").Funcs(fm).ParseGlob("variable/*"))
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
		Vivo     bool
	}
	ps := []person{
		{"Andrea", "Mazzanti", true},
		{"Stefano", "Mazzanti", true},
		{"Giuseppa", "Volpi", true},
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer nf.Close()

	err = tplnew1.ExecuteTemplate(nf, "tpl_customtypewithfunc.gohtml", ps)
	if err != nil {
		log.Fatalln(err)
	}
}
