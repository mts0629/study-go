package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of
// method signatures
type geometry interface {
	area() float64
	perim() float64
}

// Types of geometric shapes
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// Implementation of geometry on rects
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Implementation of geometry on circles
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Refer methods by an interface type
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	fmt.Println("r:")
	measure(r)
	fmt.Println("c:")
	measure(c)
}
