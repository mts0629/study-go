package main

import (
	"bufio"
	"fmt"
	"os"
)

// Error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Write a string to a file
	d1 := []byte("hello\ngo\n")
	// Permission: o+rw g+r o+r
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Create a new file
	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close() // Close a file at the end of the block

	// Write raw bytes: "some\n"
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Write a string: "writes\n"
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// bufio packge: buffered reader I/O
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n") // Write "buffered\n"
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush() // Flush all buffered operation
}
