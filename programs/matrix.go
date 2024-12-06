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

	return &Mat{
		rows: rows,
		cols: cols,
		data: data,
	}
}

func InitMat(data [][]float64) *Mat {
	rows := len(data)
	cols := 0
	for i := 0; i < rows; i++ {
		if len(data[i]) > cols {
			cols = len(data[i])
		}
	}

	mat := NewMat(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			mat.data[i][j] = data[i][j]
		}
	}

	return mat
}

func main() {
	a := InitMat([][]float64{{1, 2, 3}, {4, 5, 6}})
	fmt.Println(a.data)

	b := InitMat([][]float64{{8, 9, 10}, {11, 12, 13}})
	fmt.Println(b.data)
}
