package main

import "fmt"

func main() {
	myNumFunc := newCounter()

	fmt.Println(myNumFunc()) // 1
	fmt.Println(myNumFunc()) // 2
	fmt.Println(myNumFunc()) // 3
}

func newCounter() func() int {
	var num int

	return func() int {
		num++
		return num
	}
}
