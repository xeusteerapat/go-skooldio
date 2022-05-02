package main

import "fmt"

func main() {
	// map is key-value, zero value is "nil"
	var nilMap map[string]string

	if nilMap == nil {
		fmt.Println("nil map")
	}

	myMap := make(map[string]string)
	myMap["firstname"] = "Teerapat"
	myMap["lastname"] = "Prommarak"

	fmt.Println(myMap)

	// or create map without make keyword but is not nil zero value
	myMap2 := map[string]string{}

	if myMap2 != nil {
		fmt.Println("R U Nil, No I don't")
	}

	// access to value of map
	myFirstname := myMap["firstname"]

	fmt.Println(myFirstname) // Teerapat
}
