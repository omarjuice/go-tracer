package main

import (
	"testing"
)

func TestNewMatrix(t *testing.T) {
	m := NewMatrix(4, 4)
	var val float64
	expected := 0.0
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			val = m.Get(i, j)
			if val != expected {
				t.Errorf("NewMatrix: %f should be %f", val, expected)
			}
		}
	}
}

func TestMultiplyMatrix(t *testing.T) {
	m1 := Matrix(
		[][]float64{
			[]float64{1, 2, 3, 4},
			[]float64{5, 6, 7, 8},
			[]float64{9, 8, 7, 6},
			[]float64{5, 4, 3, 2},
		},
	)
	m2 := Matrix(
		[][]float64{
			[]float64{-2, 1, 2, 3},
			[]float64{3, 2, 1, -1},
			[]float64{4, 3, 6, 5},
			[]float64{1, 2, 7, 8},
		},
	)

	expected := Matrix(
		[][]float64{
			[]float64{20, 22, 50, 48},
			[]float64{44, 54, 114, 108},
			[]float64{40, 58, 110, 102},
			[]float64{16, 26, 46, 42},
		},
	)

	result := m1.MulMatrix(&m2)
	pass := result.Equals(&expected)

	if !pass {
		t.Errorf("MulMatrix: expected %v to be %v", result, expected)
	}

}
