package main

import "fmt"

type Book struct {
	Name   string
	Author string
}

func String(book Book) string {
	return book.Name + " by " + book.Author
}

func (book *Book) SetName(name string) {
	book.Name = name
}

func main() {
	book := Book{
		Name:   "Harry Potter",
		Author: "J. K. Rowling",
	}

	fmt.Println(String(book))

	// set name for book
	myBook := Book{
		Name:   "Test",
		Author: "Roster Armour",
	}

	myBook.SetName("Higher")

	fmt.Println(myBook) // {Higher Roster Armour}
}
