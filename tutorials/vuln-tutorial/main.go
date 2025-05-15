package main

import (
	"fmt"
	"golang.org/x/text/language"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		tag, err := language.Parse(arg)
		if err != nil {
			fmt.Printf("%s: error: %v\n", arg, err)
		} else if tag == language.Und {
			fmt.Printf("%s: undefined\n", arg)
		} else {
			// Prints a valid tag for each given BCP 47 string
			fmt.Printf("%s: tag %s\n", arg, tag)
		}
	}
}
