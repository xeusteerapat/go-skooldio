package main

import "fmt"

func main() {
	str := HOFGreeting(func() string {
		return "Teerapat"
	})

	fmt.Println(str)
}

func HOFGreeting(nameFn func() string) string {
	return fmt.Sprintf("Hello, %s", nameFn())
}
