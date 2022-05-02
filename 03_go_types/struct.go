package main

import "fmt"

type Book struct {
	name   string
	author string
}

type Rectangle struct {
	w, l float64
}

func main() {
	var designPatternsBook = Book{
		name:   "Design Patterns",
		author: "Gang of Four",
	}

	fmt.Println(designPatternsBook) // {Design Patterns Gang of Four}

	rec := Rectangle{
		w: 4.5,
		l: 6.6,
	}

	// we can mutate value by assigning
	rec.l = 7.8

	fmt.Println(rec) // {4.5 7.8}
}
