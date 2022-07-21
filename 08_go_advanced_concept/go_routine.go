package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go doSomething()
	go doSomething()
	go doSomething()
	time.Sleep(120 * time.Millisecond)
	fmt.Println(time.Since(start))
	example()
}

func doSomething() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Something func")
}

/*
Something func
Something func
Something func
122.440736ms
*/

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second) // -+--+--+--+--
}
