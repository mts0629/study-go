package main

import (
	"fmt"
	"os"
)

func write_file() {
	// Create/open a file
	f, err := os.Create("./work/test.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create file")
	}

	defer f.Close() // Defer the execution of Close()

	str := "Hello, world!\n"
	data := []byte(str)

	// Write bytes to a file
	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to write file")
	}

	fmt.Printf("Write %d bytes to \"%s\"\n", count, f.Name())
}

func read_file() {
	// Open a file
	f, err := os.Open("./work/test.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to open file")
	}

	defer f.Close()

	// Read from the file
	b := make([]byte, 16) // Allocate buffer
	n, err := f.Read(b)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to read file")
	}

	fmt.Printf("Read %d bytes from \"%s\"\n", n, f.Name())
	fmt.Printf("%s\n", string(b[:n]))
}

func main() {
	write_file()

	read_file()
}
