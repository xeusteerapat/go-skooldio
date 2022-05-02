package main

import "fmt"

func main() {
	// slice is more flexible, it can mutable and its zero value is "nil"
	var nums []int

	nums = make([]int, 4) // allowcate memory to slice with 4 elements

	nums[0] = 1
	nums[1] = 2

	fmt.Println(nums) // [1 2 0 0]

	nums = append(nums, 10) // [1 2 0 0 10]
	fmt.Println(nums)

	// behind the scene of slice is array, there are 4 properties of slice
	// data type, pointer, len, capacities
	// we can also create slice with this syntax below
	// nums = make([]int, 4, 6)
	// Go will create an array allowcate to store int for 4 element with zero value "0"
	// and its capacity is 6
	nums2 := make([]int, 4, 6)

	fmt.Println(nums2)      // [0 0 0 0]
	fmt.Println(len(nums2)) // 4
	fmt.Println(cap(nums2)) // 6
}
