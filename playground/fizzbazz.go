package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(x int) string {
	if x == 0 {
		return "0"
	}

	out := ""

	if x%3 == 0 {
		out += "Fizz"
	}
	if x%5 == 0 {
		out += "Buzz"
	}

	if out == "" {
		out = strconv.Itoa(x)
	}

	return out
}

func main() {
	var x int

	fmt.Print("Input an integer: ")
	fmt.Scan(&x)

	fmt.Println(fizzBuzz(x))
}
