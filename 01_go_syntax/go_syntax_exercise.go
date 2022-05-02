package main

import (
	"fmt"
	"strconv"
)

// greeting("Pallat") should return "Hello, Pallat"
func greeting(name string) string {
	return "Hello, " + name
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return "Hello, " + name + "." + " You are " + strconv.Itoa(int(age)) + " years old."
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return "Hello, " + name + "." + " You are " + strconv.Itoa(int(age)) + " years old" + " and your favorite drink is " + drink
}

func main() {
	a := greeting("Pallat")
	b := greetingWithAge("Pallat", 30)
	c := greetingWithAgeAndDrink("Pallat", 30, "Cola")

	fmt.Println(a, b, c)
}
