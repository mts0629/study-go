package main

import (
	"fmt"
	"testing"
)

// Function to be tested, integer minimum
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Test: MIN(2, -2)
// Function are named as "Test***"
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// Report test failures and continue tests
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Table-driven style
func TestIntMinTableDriven(t *testing.T) {
	// Test table, pairs of parameters and expected value
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// Loop for table entries
	for _, tt := range tests {
		// Name a test case from parameters
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// Run sub-tests
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Benchmarking
// Functions are named as "Benchmark***"
func BenchmarkIntMin(b *testing.B) {
	// The function will be executed many times
	// to estimate reasonable runtime
	for b.Loop() {
		IntMin(1, 2)
	}
}
