package main

import (
	"fmt"
	"slices"
)

func main() {
	// Sort strings
	strs := []string{"c", "a", "b"}
	slices.Sort(strs)
	fmt.Println("Strings:", strs)

	// Sort integers
	ints := []int{7, 2, 4}
	slices.Sort(ints)
	fmt.Println("Ints:", ints)

	// Check sorted
	s := slices.IsSorted(ints)
	fmt.Println("Sorted: ", s)
}
