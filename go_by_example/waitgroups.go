package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// WaitGroup
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Increment the WaitGroup counter

		// Closure to tell finish of worker
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Sync all workers
	wg.Wait()
}
