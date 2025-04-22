package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Set an environment variable: "FOO"=1
	os.Setenv("FOO", "1")

	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR")) // Blank if not defined

	fmt.Println()
	// Get names of current environment variables
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2) // Split key/value pairs
		fmt.Println(pair[0])              // Print keys
	}
}
