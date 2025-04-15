package main

import (
	"flag"
	"fmt"
)

func main() {
	// Declare a command line flag: `-word`
	// with a default value: "foo" and a short description: "a string"
	wordPtr := flag.String("word", "foo", "a string")

	// Declare an integer flag `-numb` and a boolean flag `-fork`
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	// Declare an option `-svar` using an existing variable
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string")

	// Execute command line parsing (`-h`/`--help` show helps)
	flag.Parse()

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)

	// Positional arguments after flags
	fmt.Println("tail:", flag.Args())
}
