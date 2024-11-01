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

func main() {
	for _, i := range []int{1, 2, 3, 4, 5} {
		if r, e := f(i); e != nil { // Inline error check
			fmt.Println("[Error] f failed:", e)
		} else {
			fmt.Println(r)
		}
	}
}
