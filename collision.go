package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Point struct
type Point struct {
	x float64 // Coordinate (x, y)
	y float64
}

// Get a distance between 2 points
func GetDist(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y

	return math.Sqrt(dx*dx + dy*dy)
}

// Ball struct
type Ball struct {
	center Point   // Center
	radius float64 // Radius
}

// Method with pointer receiver
// Move a ball
func (b *Ball) Move(dx, dy float64) {
	b.center.x += dx
	b.center.y += dy
}

// Check collision of 2 balls
func IsCollided(b1, b2 Ball) bool {
	if GetDist(b1.center, b2.center) <= (b1.radius + b2.radius) {
		return true
	}
	return false
}

func main() {
	a := Ball{Point{0, 0}, 1}
	b := Ball{Point{2, 2}, 1}

	for {
		// Random walk of 2 balls
		dx := (rand.Float64()*2 - 1.0) * 0.1
		dy := (rand.Float64()*2 - 1.0) * 0.1
		a.Move(dx, dy)

		dx = (rand.Float64()*2 - 1.0) * 0.1
		dy = (rand.Float64()*2 - 1.0) * 0.1
		b.Move(dx, dy)

		// Print position
		fmt.Printf("\ra=(%.1f, %.1f), b=(%.1f, %.1f) ", a.center.x, a.center.y, b.center.x, b.center.y)

		// If 2 balls are collided, finish
		if IsCollided(a, b) {
			fmt.Print("Collide!")
			break
		}
	}

	fmt.Println("")
}
