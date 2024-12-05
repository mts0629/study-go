package main

import "fmt"

type Mat struct {
	rows int
	cols int
	data [][]float64
}

// Constructor
func NewMat(rows int, cols int) *Mat {
	data := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		data[i] = make([]float64, cols)
	}

	m := new(Mat)
	m.rows = rows
	m.cols = cols
	m.data = data

	return m
}

func main() {
	mat := NewMat(2, 3)

	fmt.Println(mat.data)
}
