package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	// Goroutine takes 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Await c1
	select {
	case res := <-c1:
		fmt.Println(res)
	// Timeout of 1 second for c1
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1 sec.") // c1 is timeout
	}

	c2 := make(chan string, 1)
	// Goroutine takes 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// Await c2
	select {
	case res := <-c2:
		fmt.Println(res) // c2 is passed
	// Timeout of 3 seconds for c2
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3 sec.")
	}
}
