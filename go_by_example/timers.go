package main

import (
	"fmt"
	"time"
)

func main() {
	// Timer will wait 2 secconds
	timer1 := time.NewTimer(2 * time.Second)

	// Block until timer channel C sends a value
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	// Cancel the timer
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Sleep more time than timer2 durations,
	// but there's no notifications
	// -> timer2 has stopped
	time.Sleep(2 * time.Second)
}
