package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Read a file's entire contents
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Open a file
	f, err := os.Open("/tmp/dat")
	check(err)

	// Read 5 bytes from the file
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek a reading position by 6 bytes from the start of the file
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	// And read 2 bytes
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// Seek from the curent position/end of the file
	_, err = f.Seek(2, io.SeekCurrent)
	check(err)
	_, err = f.Seek(-4, io.SeekEnd)
	check(err)

	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	// io package: Read data at least 2 bytes
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Rewind a reading position
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// bufio package: buffered reader I/O
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// Close the file
	f.Close()
}
