package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, digit int) float64 {
	sq_x := 0.0
	radix := 1.0

	for i := 0; i < digit; i++ {
		var prev_sq_x float64

		for sq_x*sq_x < x {
			prev_sq_x = sq_x
			sq_x += radix
		}
		sq_x = prev_sq_x

		radix /= 10
	}

	return sq_x
}

func main() {
	digit := 12

	for _, v := range []float64{1, 2, 3, 4, 5} {
		fmt.Printf("Sqrt(%1.f) = %.15f (user): %.15f (math)\n", v, Sqrt(v, digit), math.Sqrt(v))
	}
}
