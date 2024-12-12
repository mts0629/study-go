package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send a value to each channel
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// Await two values received from the channels
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received", msg1)
		case msg2 := <-ch2:
			fmt.Println("received", msg2)
		}
	}

	ch := make(chan string, 2)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "two"
	}()

	// Await only one value received firstly
	select {
	case msg := <-ch:
		fmt.Println("received", msg)
	}
}
