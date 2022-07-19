package main

import (
	"fmt"
)

func main() {
	whatType(99)
	whatType("Hello")
}

func whatType(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("This is a string %s\n", v)
	case float64:
		fmt.Printf("This is a floating point %f\n", v)
	case int:
		fmt.Printf("This is a int %d\n", v)
	default:
		fmt.Println("Unknown Type")
	}
}
