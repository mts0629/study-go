package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") // Never printed

	// Exit this program with error code 3
	os.Exit(3)
}
