package main

import (
	"fmt"
	"time"
)

func main() {
	// Requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// limiter channel, receive a value every 200[ms]
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter                               // Blocking
		fmt.Println("request", req, time.Now()) // Timestamp every 200[ms]
	}

	// burstyLimiter channel, allow bursts of up to 3 events
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		// First 3 are received immediately
		// and remaining 2 are received every 200[ms]
		fmt.Println("request", req, time.Now())
	}
}
