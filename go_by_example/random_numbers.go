package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// Random integer n, 0 <= n < 100
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// Random float, 0.0 <= n < 1.0
	fmt.Println(rand.Float64())

	// Ranged by 5.0 <= n < 10.0
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Create a new PCG (Permuted congruential Generator) source
	s2 := rand.NewPCG(42, 1024)
	// And pass it to a generator
	r2 := rand.New(s2)
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	// Same seeds generate the same numbers
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
