package main

import "fmt"

func main() {
	myNum := 4
	double(&myNum)

	fmt.Println(myNum)

	nyNum2 := 2
	myTrippleNum := tripple(nyNum2)

	fmt.Println(myTrippleNum)
}

func double(num *int) {
	*num = *num * 2
}

func tripple(num int) int {
	return num * 3
}
