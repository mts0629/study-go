package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Error check
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Create a temporary file
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name()) // The file is named randomly

	defer os.Remove(f.Name()) // It's good to remove it explicitly

	// Data writing
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Create a temporary directory
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// Create a file in the temporary directory
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
