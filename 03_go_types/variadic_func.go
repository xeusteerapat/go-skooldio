package main

import "fmt"

// this function can receive multiple strings
// Go will convert multiple strings to slice of string
func variadic(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println("=============")
}

func main() {
	variadic("Teerapat", "Prommarak", "Xeus", "Cool")

	// pass slice to variadic
	s := []string{"I", "am", "Legend"}
	variadic(s...)
}
