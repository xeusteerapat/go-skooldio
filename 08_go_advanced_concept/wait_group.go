package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// add wait group
	wg := &sync.WaitGroup{}

	wg.Add(3)

	start := time.Now()
	go doSomething2(wg)
	go doSomething2(wg)
	go doSomething2(wg)
	// time.Sleep(120 * time.Millisecond)
	wg.Wait()
	fmt.Println(time.Since(start))
}

func doSomething2(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Something func")
	wg.Done()
}

/*
Something func
Something func
Something func
122.440736ms
*/
