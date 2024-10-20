package main

import "fmt"

// Calculate n-th number of the Fibonacci sequence
func fib(n int) int {
	if n < 2 {
		return n
	}

	// Recursive call
	return fib(n-2) + fib(n-1)
}

func main() {
	nums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range nums {
		fmt.Println(fib(n))
	}
}
