package main

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"
)

type Name struct {
	First, Last string
}

type Book struct {
	Title       string
	PageCount   int
	ISBN        string
	Authors     []Name
	Publisher   string
	PublishDate time.Time
}

func main() {
	books := []Book{
		{
			Title:       "Dal Big Bang ai buchi neri",
			PageCount:   3000,
			ISBN:        "978-88-17-10594-1",
			Authors:     []Name{{"Stephen", "Hawking"}},
			Publisher:   "BUR",
			PublishDate: time.Date(206, time.July, 0, 0, 0, 0, 0, time.UTC),
		},
	}
	file, err := os.Create("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	enc := gob.NewEncoder(file)
	if err := enc.Encode(books); err != nil {
		fmt.Println(err)
	}
	file, err = os.Create("book.dat")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	var loadedBooks []Book
	dec := gob.NewDecoder(file)
	if err := dec.Decode(&loadedBooks); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(loadedBooks)
}
