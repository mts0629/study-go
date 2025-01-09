package main

import (
	"fmt"
	"time"
)

// Worker goroutine
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	// Worker pool
	const numJobs = 5
	jobs := make(chan int, numJobs)    // Senders
	results := make(chan int, numJobs) // Receivers

	// Start up 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs -> 5 jobs in 3 workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Notify that's all

	// Collect all results
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
