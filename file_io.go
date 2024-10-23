package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./work/test.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to create file")
	}

	str := "Hello, world!\n"
	data := []byte(str)

	count, err := f.Write(data)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Failed to write file")
	}

	fmt.Printf("Write %d bytes\n", count)

	defer f.Close()
}
