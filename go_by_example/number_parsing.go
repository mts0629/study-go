package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Parse a float in 64-bit precision
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Parse an integer in 64-bit, base from the string
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// Hex format
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Parse an unsigned integer
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Convert a string to a base-10 integer
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// Error
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
