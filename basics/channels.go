package main

import "fmt"

func main() {
	// New channel which conveys a string
	messages := make(chan string)

	// Send a string into the channel in a goroutine
	go func() { messages <- "put" }()

	// Receive a string from the channel
	msg := <-messages
	fmt.Println(msg)

	// Buffered channel which accepts 2 values
	// without concurrent receive
	buf := make(chan string, 2)
	buf <- "hello"
	buf <- "world"

	// Values are queued
	fmt.Println(<-buf)
	fmt.Println(<-buf)
}
