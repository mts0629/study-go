package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler
func hello(w http.ResponseWriter, req *http.Request) {
	// Context carries information across APIs/goroutines
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second): // Do some task
		// And then reply "hello\n"
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		// If the work is cancelled,
		// carry an cancellation signal to other goroutines by a context
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		// And print why the task is cancelled
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// Register the handler and start a server
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
