package main

import "fmt"

func main() {
	var a interface{}

	a = 10
	fmt.Printf("%T %v\n", a, a) // int 10

	a = "Ten"
	fmt.Printf("%T %v\n", a, a) // string Ten

	a = true
	fmt.Printf("%T %v\n", a, a) // bool true

	a = func() string { return "Hi" }
	fmt.Printf("%T %v\n", a, a) // func() string 0x108ac00

	fmt.Println()
}
