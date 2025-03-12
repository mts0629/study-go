package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	// defer, call closeFile(f) at the end of main()
	defer closeFile(f)

	writeFile(f)
	// closeFile(f) will be executed here
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil { // Error should be checked
		panic(err)
	}
}
