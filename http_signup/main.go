package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template

// DB Users (in memory)
var dbUsers = map[string]user{} //UUID -> users
//DB Sessions (in memory)
var dbSessions = map[string]string{} //session-ID -> UUID

func init() {
	tpl = template.Must(template.ParseGlob("http_signup/templates/*"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["mazzantia@hotmail.com"] = user{
		UserName: "mazzantia@hotmail.com",
		Password: bs,
		First:    "Andrea",
		Last:     "Mazzanti",
	}

}

func main() {
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil) //using GO default handler
}

func index(w http.ResponseWriter, req *http.Request) {
	_, u := getUser(req, w)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	_, u := getUser(req, w)
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	//check if already logged-in
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}
	//get form values
	if req.Method == http.MethodPost {

		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		//username already there
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username not available", http.StatusForbidden)
			return
		}
		//create a session and store it on the cookie
		id, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			Secure:   false,
			HttpOnly: true,
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		//store user to DB
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		u := user{un, bs, f, l}
		dbUsers[un] = u
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func getUser(req *http.Request, w http.ResponseWriter) (*http.Cookie, user) {
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
	return c, u
}

func alreadyLoggedIn(req *http.Request) bool {
	//get cookie
	c, err := req.Cookie("session-id")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok

}
func login(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		id, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:     "session-id",
			Value:    id.String(),
			Secure:   false,
			HttpOnly: true,
		}

		http.SetCookie(w, c)
		dbSessions[c.Value] = un
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	c, _ := req.Cookie("session-id")
	// delete the session
	delete(dbSessions, c.Value)
	// remove the cookie
	c = &http.Cookie{
		Name:   "session-id",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
