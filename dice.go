package main

import (
	"flag"
	"fmt"
	"math/rand/v2"
)

func main() {
	// Command line options
	var (
		s = flag.Int("s", 6, "number of sides")
		n = flag.Int("n", 1, "number of dice rolls")
	)

	flag.Parse()

	// Dice roll
	for i := 0; i < *n; i++ {
		fmt.Print(rand.IntN(*s))

		if i < *n-1 {
			fmt.Print(",")
		}
	}

	fmt.Println()
}
