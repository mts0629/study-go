package main

import (
	"fmt"
	"os"
)

func main() {
	// Get raw command line arguments
	// The first value is the path to the program
	argsWithProg := os.Args
	fmt.Println(argsWithProg)

	// Skip program
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithoutProg)

	// Get the third argument
	arg := os.Args[3]
	fmt.Println(arg)
}
