package main

import "fmt"

func main() {
	defer fmt.Println("# defer")
	printDefer()
	trickyDefer()

	fmt.Println("Print first") // Print first and then # defer
}

func printDefer() {
	defer fmt.Println("defer #1")
	defer fmt.Println("defer #2")
	defer fmt.Println("defer #3")
}

/*
defer #3
defer #2
defer #1
Print first
# defer
*/

func trickyDefer() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

/*
3
3
3
*/
