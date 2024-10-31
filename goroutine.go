package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function
func worker(msg string, wg *sync.WaitGroup) {
	defer wg.Done() // Notify this the worker is done on the last

	for i := 0; i < 5; i++ {
		fmt.Println(msg, ":", i)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	// WaitGroup for synchronization of goroutines
	var wg sync.WaitGroup

	// Goroutine, lightweight thread
	wg.Add(1)
	go worker("Goroutine #1", &wg)
	wg.Add(1)
	go worker("Goroutine #2", &wg)

	// Wait finish of goroutines
	wg.Wait()
}
