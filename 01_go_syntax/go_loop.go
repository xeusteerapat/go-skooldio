package main

import (
	"fmt"
	"math"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func main() {
	ans := IsPrime(3)
	fmt.Println(ans)
}
