package main

import (
	"fmt"
	"strings"
)

func main() {
	result := wordCount("Apple Banana Apple Banana apple")

	fmt.Println(result) // map[Apple:2 Banana:2 apple:1]
}

func wordCount(s string) map[string]int {
	splitWords := strings.Split(s, " ")

	result := map[string]int{}

	for i := 0; i < len(splitWords); i++ {
		result[splitWords[i]] = result[splitWords[i]] + 1
	}

	return result
}
