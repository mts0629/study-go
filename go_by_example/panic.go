package main

import "os"

func main() {
	// Panic, unexpected error
	// panic("a problem")

	// Panic on failure of file creation
	_, err := os.Create("/file")
	if err != nil {
		panic(err)
	}
}
