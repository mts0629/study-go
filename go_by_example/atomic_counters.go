package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// Atomic counter accessed by multiple goroutines
	var ops atomic.Uint64

	var wg sync.WaitGroup

	// Start 50 goroutines
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// Increment counter 1000 times in one goroutine
		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1) // Increment atomically
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops: ", ops.Load()) // Read atomically
}
