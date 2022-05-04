package main

import (
	"fmt"

	"github.com/xeusteerapat/demo/book"
)

func main() {
	newBook := book.New()

	fmt.Println(newBook) // empty struct { }
}
