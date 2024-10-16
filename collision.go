package main

import (
	"fmt"
	"math"
)

// Ball struct
type Ball struct {
	x float64 // Center
	y float64
	radius float64 // Radius
}

// Method
func (b1 Ball) Collide(b2 Ball) bool {
	dx := b1.x - b2.x
	dy := b1.y - b2.y

	if math.Sqrt(dx * dx + dy * dy) <= (b1.radius + b2.radius) {
		return true
	}

	return false
}

// Method with pointer
func (b *Ball) Move(dx, dy float64) {
	b.x += dx
	b.y += dy
}

func main() {
	a := Ball{1, 1, 1}
	b := Ball{2, 2, 1}

	fmt.Println("a=(", a.x, ",", a.y, "),", a.Collide(b))

	a.Move(1, 1)
	fmt.Println("a=(", a.x, ",", a.y, "),", a.Collide(b))
}
