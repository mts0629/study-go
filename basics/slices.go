package main

import (
	"fmt"
	"slices"
)

func main() {
	// Uninitialized slice
	var s []string
	fmt.Println("s =", s, "len =", len(s), "cap =", cap(s))
	if s == nil {
		fmt.Println("s is nil")
	}

	// Zero-initialized slice
	s = make([]string, 3)
	fmt.Println("s =", s, "len =", len(s), "cap =", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("s =", s, "len =", len(s))
	fmt.Println("s[2] =", s[2])

	// Appending elements
	s = append(s, "d")
	fmt.Println("s =", s, "len =", len(s))
	s = append(s, "e", "f")
	fmt.Println("s =", s, "len =", len(s))

	// Copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c =", c, "len =", len(c))

	// "slice" operator
	l := s[2:5]
	fmt.Println("l =", l, "len =", len(l))
	l = s[:4] // s[4] is excluded
	fmt.Println("l =", l, "len =", len(l))
	l = s[3:] // s[3] is included
	fmt.Println("l =", l, "len =", len(l))

	// Initialized slice
	t := []int{1, 2, 3}
	fmt.Println("t =", t, "len =", len(t))

	// Check equality of 2 slices
	t2 := []int{1, 2, 3}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// Multi-dimensional slice
	mat := make([][]int, 3) // First dimension
	for i := 0; i < 3; i++ {
		second_dim := i + 1
		mat[i] = make([]int, second_dim) // Second dimension
		for j := 0; j < len(mat[i]); j++ {
			mat[i][j] = j + second_dim
		}
	}
	fmt.Println("mat =", mat) // Can be "jagged"
}
