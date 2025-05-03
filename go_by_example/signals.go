package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create a channel to receive a signal
	sigs := make(chan os.Signal, 1)

	// Notify when receiving SIGINT and SIGTERM
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	// Goroutine: print a signal
	// and notify a finish of the program
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done // Blocking
	fmt.Println("exiting")
}
