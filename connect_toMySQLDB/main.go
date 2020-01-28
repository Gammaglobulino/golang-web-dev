package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:Pacifista001!@tcp(localhost:3306)/population?charset=utf8")
	check(err)
	defer db.Close()
	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/peoples", peoples)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)

}
func peoples(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query("SELECT NAME FROM persons")
	check(err)
	var s, name string
	s = "RETRIEVED RECORDS:\n"
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintf(w, s)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err = io.WriteString(w, "DB Connection completed")
	check(err)
}
func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
