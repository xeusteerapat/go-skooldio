package main

import "fmt"

func isOdd(n int) bool {
	if n%2 == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	beanloo := isOdd(4)
	fmt.Println(beanloo)
}
