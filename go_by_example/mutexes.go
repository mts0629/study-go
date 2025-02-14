package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex // Mutex
	counters map[string]int
}

// Increment counter
// Mutex should be passed by pointer
func (c *Container) inc(name string) {
	c.mu.Lock()         // Lock
	defer c.mu.Unlock() // Unlock
	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Goroutine, increment named counter in the loop
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Run goroutines concurrently
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	wg.Wait()
	fmt.Println(c.counters)
}
