package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:Pacifista001!@localhost/gammawar?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the DB")
}

type Employee struct {
	id      int
	name    string
	rank    int
	address string
	salary  float64
	bday    time.Time
}

func main() {
	//test with http://localhost:8080/employees/show?name=Andrea
	defer db.Close()
	http.HandleFunc("/employees", employeesIndex)
	http.HandleFunc("/employees/show", employeeShow)
	http.ListenAndServe(":8080", nil)
}

func employeeShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { //do only GET
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	name := r.FormValue("name")
	fmt.Println(name)
	if name == "" { //do only GET
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	row := db.QueryRow("SELECT * FROM employees WHERE employees.name like $1", name)
	emp := Employee{}
	err := row.Scan(&emp.id, &emp.name, &emp.rank, &emp.address, &emp.salary, &emp.bday)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, r)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), 500)
		return

	}
	fmt.Fprintf(w, "%d %s,%d,%s,%f,%s\n", emp.id, emp.name, emp.rank, emp.address, emp.salary, emp.bday)
}

func employeesIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" { //do only GET
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("select * from employees;")
	if err != nil {
		fmt.Println(err)
		http.Error(w, http.StatusText(500), 500) //internal server error
		return
	}
	defer rows.Close()

	emps := make([]Employee, 0)
	for rows.Next() {
		emp := Employee{}
		err := rows.Scan(&emp.id, &emp.name, &emp.rank, &emp.address, &emp.salary, &emp.bday)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
		}
		emps = append(emps, emp)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, emp := range emps {
		fmt.Fprintf(w, "%d %s,%d,%s,%f,%s\n", emp.id, emp.name, emp.rank, emp.address, emp.salary, emp.bday)
	}

}
