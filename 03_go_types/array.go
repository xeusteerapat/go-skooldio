package main

import "fmt"

func main() {
	// array is immutable, cannot increase/decrease size
	var fourNums [4]int
	fourNums[0] = 1
	fourNums[1] = 2

	fmt.Print(fourNums) // [1 2 0 0]
}
