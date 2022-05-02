package main

import "fmt"

func main() {
	var s string
	var p *string

	p = &s // bring address of s (0xc000096210) and put it to p

	*p = "I am P" // so s = "I am P"

	fmt.Println("s ->", s)
	fmt.Println("p ->", p)
}
