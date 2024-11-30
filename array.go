package main

import "fmt"

func main() {
	// Create an array
	var a [5]int
	fmt.Println("a =", a, "len =", len(a))

	// Get/set an element
	a[3] = 3
	fmt.Println("a =", a, "a[3] =", a[3])

	// Initialization
	b := [...]int{1, 2, 3}
	fmt.Println("b = ", b)

	// Initialization with length (zero-fill)
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	// Initialization with an index
	d := [...]int{1, 3: 4, 5}
	fmt.Println("d = ", d)

	// Multi-dimensional array
	m1 := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println("m1 =", m1)

	m2 := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("m2 =", m2)

	var m3 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	fmt.Println("m3 =", m3)
}
