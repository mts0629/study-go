package main

import (
	"fmt"
	"time"
)

// Worker goroutine with channel
func worker(finish chan bool) {
	fmt.Print("wait...")
	time.Sleep(time.Second)
	fmt.Println("finished")

	// Notify finish of the work
	finish <- true
}

// Only accept a channel for sending value
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// First channel for sending value, second channel for receiving value
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg + "pong"
}

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

	finish := make(chan bool, 1)
	go worker(finish)

	// Block until receiving a notification
	<-finish

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
