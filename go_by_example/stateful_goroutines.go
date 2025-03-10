package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Message and corresponding replies
type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	// Channels for read/write request
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Goroutine which owns the state
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				// Read response is a value
				read.resp <- state[read.key]
			case write := <-writes:
				// Write response is success of writing (true)
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Read
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1) // Add counts
				time.Sleep(time.Millisecond)  // Wait 1 ms
			}
		}()
	}

	// Write
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1) // Add counts
				time.Sleep(time.Millisecond)   // Wait 1 ms
			}
		}()
	}

	time.Sleep(time.Second)

	// Number of completed operations
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
