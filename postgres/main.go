package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	db, err := sql.Open("postgres", "postgres://postgres:Pacifista001!@localhost/gammawar?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the DB")
	rows, err := db.Query("select * from employees;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	books := make([]struct {
		id     int
		name   string
		rank   int
		adress string
		salary float64
		bday   time.Time
	}, 0)
	for rows.Next() {
		bk := struct {
			id     int
			name   string
			rank   int
			adress string
			salary float64
			bday   time.Time
		}{}
		err := rows.Scan(&bk.id, &bk.name, &bk.rank, &bk.adress, &bk.salary, &bk.bday)
		if err != nil {
			panic(err)
		}
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	for _, bk := range books {
		fmt.Printf("%d %s,%d,%s,%f,%s\n", bk.id, bk.name, bk.rank, bk.adress, bk.salary, bk.bday)
	}

}
