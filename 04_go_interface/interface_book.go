package main

import "fmt"

type Stringer interface {
	String() string
}

type Book struct {
	Name   string
	Author string
}

// create pointer receiver for book type
func (book Book) String() string {
	return book.Name + " by " + book.Author
}

func main() {
	designPatternsBook := Book{
		Name:   "The Design Patterns",
		Author: "Gang of Four",
	}

	fmt.Println(designPatternsBook.String())
}
