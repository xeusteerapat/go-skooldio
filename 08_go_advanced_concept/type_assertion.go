package main

import "fmt"

func main() {
	var a interface{}
	a = 10
	fmt.Printf("%T %v\n", a, a)

	if num, ok := a.(int); ok { // check if type assertion is ok
		fmt.Println(num)
	}
}
