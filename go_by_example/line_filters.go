package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Wrap stdin with a buffered scanner
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() { // Advance the scanner to the next token
		// Get the current token and capiterize it
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Error check
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
