package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Point struct
type Point struct {
	x float64 // Coordinate x
	y float64 // Coordinate y
}

// Get a distance between 2 points
func GetDist(p1, p2 Point) float64 {
	dx := p1.x - p2.x
	dy := p1.y - p2.y

	return math.Sqrt(dx*dx + dy*dy)
}

// Ball struct
type Ball struct {
	Point          // Coordinate, embedded struct
	radius float64 // Radius
}

// Method with pointer receiver
// Move a ball
func (b *Ball) Move(dx, dy float64) {
	b.x += dx
	b.y += dy
}

// Methos with value receiver
// Print a position
func (b Ball) GetPos() (float64, float64) {
	return b.x, b.y
}

// Check collision of 2 balls
func IsCollided(b1, b2 Ball) bool {
	if GetDist(b1.Point, b2.Point) <= (b1.radius + b2.radius) {
		return true
	}
	return false
}

// Get random (x, y) values in [-1, 1]
func getRandXY() (x, y float64) {
	x = rand.Float64()*2 - 1.0
	y = rand.Float64()*2 - 1.0
	return
}

func main() {
	a := Ball{Point{0, 0}, 1}
	b := Ball{Point{10, 10}, 1}

	start := time.Now()

	for {
		// Random walk of 2 balls
		a.Move(getRandXY())
		b.Move(getRandXY())

		// Print positions of the balls
		ax, ay := a.GetPos()
		bx, by := b.GetPos()
		fmt.Printf("\ra=(%.2f, %.2f), b=(%.2f, %.2f) ", ax, ay, bx, by)

		// If 2 balls are collided, finish the loop
		if IsCollided(a, b) {
			fmt.Print("Collide!")
			break
		}

		// Wait 16[ms] (-> 60[FPS])
		time.Sleep(16 * time.Millisecond)

		// If over 5[sec], end the loop
		if time.Since(start) >= (5 * time.Second) {
			break
		}
	}

	fmt.Println("") // End the line
}
