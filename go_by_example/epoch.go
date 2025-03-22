package main

import (
	"fmt"
	"time"
)

func main() {
	// Current time
	now := time.Now()
	fmt.Println(now)

	// Unix epoch
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli()) // Milliseconds
	fmt.Println(now.UnixNano())  // Nanoseconds

	// Convert seconds to time
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
