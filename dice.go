package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	// Specify number of rolls, 1 in default
	n_rolls := 1
	if len(args) == 1 {
		var err error
		n_rolls, err = strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	for i := 0; i < n_rolls; i++ {
		fmt.Print(rand.IntN(6))

		if i < n_rolls-1 {
			fmt.Print(",")
		}
	}
}
