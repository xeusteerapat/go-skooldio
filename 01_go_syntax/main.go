package main

import "fmt"

func greeting(firstname, lastname string) {
	fmt.Println("Hello,", firstname, lastname)
}

func add(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	fmt.Println("hi")
	greeting("Teerapat", "Prommarak")
	right, left := swap(4, 5)
	fmt.Println(right, left)
}
