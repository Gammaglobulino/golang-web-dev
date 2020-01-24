package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template

// DB Users (in memory)
var dbUsers = map[string]user{} //UUID -> users
//DB Sessions (in memory)
var dbSessions = map[string]string{} //session-ID -> UUID

func init() {
	tpl = template.Must(template.ParseGlob("http_sessions_and_UUID/templates/*"))
}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", base)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil) //using GO default handler
}

func base(w http.ResponseWriter, req *http.Request) {
	//read cookie
	c, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}
	//UUID is there get the user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	//process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("uuid")
		fn := req.FormValue("firstname")
		ln := req.FormValue("lastname")
		u = user{
			UserName: un,
			First:    fn,
			Last:     ln,
		}
		dbSessions[c.Value] = un
		dbUsers[un] = u
		fmt.Println(un, fn, ln)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
func bar(w http.ResponseWriter, req *http.Request) {
	//get cookie
	c, err := req.Cookie("session-id")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}
