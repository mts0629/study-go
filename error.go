package main

import (
	"errors"
	"fmt"
)

// Return error
// The last return value with type error, by convention
func f(value int) (int, error) {
	if value == 3 {
		return -1, errors.New("3 is invalid")
	}

	return value, nil // nil when there's no error
}

// Sentinel errors
var ErrDivisibleBy2 = fmt.Errorf("Divisible by 2")
var ErrDivisibleBy3 = fmt.Errorf("Divisible by 3")

func div(value int) error {
	if value%2 == 0 {
		return ErrDivisibleBy2
	} else if value%3 == 0 {
		return ErrDivisibleBy3
	}

	return nil
}

func main() {
	for i := 0; i < 5; i++ {
		if r, e := f(i); e != nil { // Inline error check
			fmt.Println("[Error] f failed:", e)
		} else {
			fmt.Println(r)
		}
	}

	fmt.Println("*****")

	for i := 1; i <= 5; i++ {
		if err := div(i); err != nil {
			if errors.Is(err, ErrDivisibleBy2) { // Check error matching
				fmt.Println("[Error] Divisible by 2")
			} else if errors.Is(err, ErrDivisibleBy3) {
				fmt.Println("[Error] Divisible by 3")
			} else {
				fmt.Printf("[Error] Unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println(i)
	}
}
