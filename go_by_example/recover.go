package main

import "fmt"

// Function which causes panic
func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		// recover must be called within a defered function
		if r := recover(); r != nil {
			// Return value is the error raised from the panic
			fmt.Println("Recovered.\nError:", r)
		}
	}()

	mayPanic()

	// Not executed
	fmt.Println("After mayPanic()")
}
