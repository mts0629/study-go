package main

import "fmt"

// Variadic function: print a sum of arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")

	total := 0
	for _, num := range nums { // nums equivalent to []int
		total += num
	}

	fmt.Println(total)
}

func main() {
	// Pass multiple args
	sum(1, 2)
	sum(1, 2, 3)

	// Pass multiple args in a slice
	num := []int{1, 2, 3, 4}
	sum(num...)
}
