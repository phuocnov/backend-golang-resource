package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		fmt.Println("Hello ", i)
		time.Sleep(500 * time.Millisecond)
		wg.Done()
	}
}

var wg = sync.WaitGroup{}

func main() {
	// say hello with go rountine
	sayHello()
	wg.Wait()
}
