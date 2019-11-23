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

func TestMatrixMulMatrix(t *testing.T) {
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

func TestMatrixMulTuple(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{1, 2, 3, 4},
			[]float64{2, 4, 4, 2},
			[]float64{8, 6, 4, 1},
			[]float64{0, 0, 0, 1},
		},
	)
	tuple := &Tuple{1, 2, 3, 1}

	expected := &Tuple{18, 24, 33, 1}

	result := m.MulTuple(tuple)

	if !result.Equals(expected) {
		t.Errorf("MatrixMulTuple: expected %v to be %v", result, expected)
	}
}

func TestIdentityMatrix(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{20, 22, 50, 48},
			[]float64{44, 54, 114, 108},
			[]float64{40, 58, 110, 102},
			[]float64{16, 26, 46, 42},
		},
	)

	if !m.Equals(m.MulMatrix(&IdentityMatrix)) {
		t.Errorf("IdentityMatrix invalid.")
	}

}
func TestTransposeMatrix(t *testing.T) {
	if !IdentityMatrix.Transpose().Equals(&IdentityMatrix) {
		t.Errorf("MatrixTranspose on IdentityMatrix failed")
	}

	m := Matrix(
		[][]float64{
			[]float64{0, 9, 3, 0},
			[]float64{9, 8, 0, 8},
			[]float64{1, 8, 5, 3},
			[]float64{0, 0, 5, 8},
		},
	)
	expected := Matrix(
		[][]float64{
			[]float64{0, 9, 1, 0},
			[]float64{9, 8, 8, 0},
			[]float64{3, 0, 5, 5},
			[]float64{0, 8, 3, 8},
		},
	)
	result := m.Transpose()

	if !result.Equals(&expected) {
		t.Errorf("MatrixTranspose: expected %v to equal %v", result, expected)

	}

}
