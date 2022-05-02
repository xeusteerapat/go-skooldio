package main

import "fmt"

func main() {
	var emptySlice []int
	if emptySlice == nil {
		fmt.Println("nil slice")
	}

	// create array and slice
	// [...] is create dynamic size of array, then let compiler count and put the exact number
	// put in [] during the compile time
	arr := [...]int{1, 3, 5, 7, 9}
	slice := arr[:]

	fmt.Printf("%d %d %p %p\n", len(slice), cap(slice), slice, &arr) // 5 5 0xc000122030 0xc000122030

	slice = append(slice, 11, 13)
	// Go will create new array, as you can see 0xc000132000 != 0xc000132000
	fmt.Printf("%d %d %p %p\n", len(slice), cap(slice), slice, &arr) // 7 10 0xc000132000 0xc000122030
}
