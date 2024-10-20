package main

import "fmt"

func sequence() func() int {
	i := 0
	// Closure, return an anonymous function
	return func() int {
		i++
		return i
	}

}

func main() {
	next := sequence()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}
