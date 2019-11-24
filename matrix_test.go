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

	result := m1.MulMatrix(m2)
	pass := result.Equals(expected)

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

	if !m.Equals(m.MulMatrix(IdentityMatrix)) {
		t.Errorf("IdentityMatrix invalid.")
	}

}
func TestTransposeMatrix(t *testing.T) {
	if !IdentityMatrix.Transpose().Equals(IdentityMatrix) {
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

	if !result.Equals(expected) {
		t.Errorf("MatrixTranspose: expected %v to equal %v", result, expected)

	}

}

func TestSubMatrix(t *testing.T) {
	m1 := Matrix(
		[][]float64{
			[]float64{1, 5, 0},
			[]float64{-3, 2, 7},
			[]float64{0, 6, -3},
		},
	)
	m1ResultSub := m1.SubMatrix(0, 2)
	m1ExpectedSub := Matrix(
		[][]float64{
			[]float64{-3, 2},
			[]float64{0, 6},
		},
	)

	if !m1ResultSub.Equals(m1ExpectedSub) {
		t.Errorf("MatrixSubMatrix: expected %v to equal %v", m1ResultSub, m1ExpectedSub)
	}

	m2 := Matrix(
		[][]float64{
			[]float64{-6, 1, 1, 6},
			[]float64{-8, 5, 8, 6},
			[]float64{-1, 0, 8, 2},
			[]float64{-7, 1, -1, 1},
		},
	)
	m2ResultSub := m2.SubMatrix(2, 1)
	m2ExpectedSub := Matrix(
		[][]float64{
			[]float64{-6, 1, 6},
			[]float64{-8, 8, 6},
			[]float64{-7, -1, 1},
		},
	)
	if !m2ExpectedSub.Equals(m2ResultSub) {
		t.Errorf("MatrixSubMatrix: expected %v to equal %v", m2ResultSub, m2ExpectedSub)
	}

}

func TestMatrixMinor(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{3, 5, 0},
			[]float64{2, -1, -7},
			[]float64{6, -1, 5},
		},
	)
	result := m.Minor(1, 0)
	expected := 25.0

	if result != expected {
		t.Errorf("MatrixMinor: expected %f to equal %f", result, expected)
	}
}

func TestMatrixCofactor(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{3, 5, 0},
			[]float64{2, -1, -7},
			[]float64{6, -1, 5},
		},
	)
	minor1 := m.Minor(0, 0)
	cofactor1 := m.Cofactor(0, 0)
	minor2 := m.Minor(1, 0)
	cofactor2 := m.Cofactor(1, 0)

	if !floatEqual(minor1, -12) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", minor1, -12.0)
	}
	if !floatEqual(cofactor1, -12) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", cofactor1, -12.0)
	}
	if !floatEqual(minor2, 25) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", minor2, 25.0)
	}
	if !floatEqual(cofactor2, -25) {
		t.Errorf("MatrixCofactor: expected %f to equal %f", cofactor2, -25.0)
	}

}

func TestMatrixDeterminant(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{1, 5},
			[]float64{-3, 2},
		},
	)

	result := m.Determinant()
	expected := 17.0

	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}

	m = Matrix(
		[][]float64{
			[]float64{1, 2, 6},
			[]float64{-5, 8, -4},
			[]float64{2, 6, 4},
		},
	)
	result = m.Determinant()
	expected = -196.0

	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}
	m = Matrix(
		[][]float64{
			[]float64{-2, -8, 3, 5},
			[]float64{-3, 1, 7, 3},
			[]float64{1, 2, -9, 6},
			[]float64{-6, 7, 7, -9},
		},
	)
	result = m.Determinant()
	expected = -4071.0
	if !floatEqual(result, expected) {
		t.Errorf("MatrixDeterminant: expected %v to equal %v", result, expected)
	}
}

func TestMatrixInverse(t *testing.T) {
	m := Matrix(
		[][]float64{
			[]float64{8, -5, 9, 2},
			[]float64{7, 5, 6, 1},
			[]float64{-6, 0, 9, 6},
			[]float64{-3, 0, -9, -4},
		},
	)
	expected := Matrix(
		[][]float64{
			[]float64{-0.15385, -0.15385, -0.28205, -0.53846},
			[]float64{-0.07692, 0.12308, 0.025641, 0.03077},
			[]float64{0.35897, 0.35897, 0.43590, 0.92308},
			[]float64{-0.69230, -0.69231, -0.76923, -1.92308},
		},
	)

	result := m.Inverse()
	if !result.Equals(expected) {

		t.Errorf("MatrixInverse: result %v does not equal %v", result, expected)

	}

	a := Matrix(
		[][]float64{
			[]float64{1, 2, 3, 4},
			[]float64{5, 6, 7, 8},
			[]float64{9, 8, 7, 6},
			[]float64{5, 4, 3, 2},
		},
	)
	b := Matrix(
		[][]float64{
			[]float64{-2, 1, 2, 3},
			[]float64{3, 2, 1, -1},
			[]float64{4, 3, 6, 5},
			[]float64{1, 2, 7, 8},
		},
	)

	c := a.MulMatrix(b)

	if !c.MulMatrix(b.Inverse()).Equals(a) {
		t.Errorf("MatrixInverse: multiply by inverse does not work.")
	}
}
