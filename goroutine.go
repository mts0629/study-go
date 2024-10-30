package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, ":", i)
	}
}

func main() {
	// Goroutine, lightweight thread
	go f("Goroutine #1")
	go f("Goroutine #2")

	time.Sleep(time.Second)
}
