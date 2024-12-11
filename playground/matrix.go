package main

import "fmt"

type Mat struct {
	rows int
	cols int
	data [][]float64
}

// Constructor
func NewMat(rows, cols int) *Mat {
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
		copy(mat.data[i], data[i])
	}

	return mat
}

var ErrUnmatchMatrixSize = fmt.Errorf("Matrix size does not match")

func (a *Mat) Add(k float64, b *Mat) (*Mat, error) {
	if a.rows != b.rows || a.cols != b.cols {
		return nil, ErrUnmatchMatrixSize
	}

	c := NewMat(a.rows, a.cols)

	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			c.data[i][j] = a.data[i][j] + k*b.data[i][j]
		}
	}

	return c, nil
}

func (a *Mat) Scale(k float64) *Mat {
	b := NewMat(a.rows, a.cols)

	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			b.data[i][j] = k * a.data[i][j]
		}
	}

	return b
}

func (a *Mat) Transpose() *Mat {
	t := NewMat(a.cols, a.rows)

	for i := 0; i < a.rows; i++ {
		for j := 0; j < a.cols; j++ {
			t.data[j][i] = a.data[i][j]
		}
	}

	return t
}

func (a *Mat) MatMul(b *Mat) (*Mat, error) {
	if a.rows != b.cols {
		return nil, ErrUnmatchMatrixSize
	}

	c := NewMat(a.rows, b.cols)

	for i := 0; i < a.rows; i++ {
		for j := 0; j < b.cols; j++ {
			mac := .0
			for k := 0; k < a.cols; k++ {
				mac += a.data[i][k] * b.data[k][j]
			}
			c.data[i][j] = mac
		}
	}

	return c, nil
}

func main() {
	a := InitMat([][]float64{{1, 2, 3}, {4, 5, 6}})
	fmt.Println("a =", a.data)

	b := InitMat([][]float64{{8, 9, 10}, {11, 12, 13}})
	fmt.Println("b =", b.data)

	sum, _ := a.Add(1, b)
	fmt.Println("a + b =", sum.data)
	sub, _ := a.Add(-1, b)
	fmt.Println("a - b =", sub.data)

	ka := a.Scale(2)
	fmt.Println("2 * a =", ka.data)

	ta := a.Transpose()
	fmt.Println("T(a) =", ta.data)

	if r, e := a.Add(1, ta); e != nil {
		fmt.Println("a + T(a) failed")
	} else {
		fmt.Println("a + T(a) =", r.data)
	}

	mul, _ := a.MatMul(ta)
	fmt.Println("a * T(a) =", mul.data)

	if r, e := a.MatMul(b); e != nil {
		fmt.Println("a * b failed")
	} else {
		fmt.Println("a * b =", r.data)
	}
}
