package main

//Matrix ...
type Matrix [][]float64

//NewMatrix creates a rowsXcols matrix
func NewMatrix(rows, cols int) Matrix {
	matrix := make([][]float64, rows, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols, cols)
	}
	return matrix
}

//Set sets a value in a matrix
func (matrix *Matrix) Set(r, c int, val float64) float64 {
	(*matrix)[r][c] = val
	return val
}

//Get gets a value in a matrix
func (matrix *Matrix) Get(r, c int) float64 {
	return (*matrix)[r][c]

}

//Equals tells if two matrices are equal
func (matrix *Matrix) Equals(other *Matrix) bool {
	m := *matrix
	o := *other
	if len(m) != len(o) {
		return false
	}
	if len(m[0]) != len(o[0]) {
		return false
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] != o[i][j] {
				return false
			}
		}
	}
	return true
}

//Size returns the height and width of the matrix
func (matrix *Matrix) Size() (int, int) {
	h := len(*matrix)
	w := 0
	if h > 0 {
		w = len((*matrix)[0])
	}
	return h, w
}

//MulMatrix multiplies two 4x4 matrices together
func (matrix *Matrix) MulMatrix(other *Matrix) *Matrix {
	m := *matrix
	o := *other
	newM := NewMatrix(4, 4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			product := m.Get(row, 0)*o.Get(0, col) + m.Get(row, 1)*o.Get(1, col) + m.Get(row, 2)*o.Get(2, col) + m.Get(row, 3)*o.Get(3, col)
			newM.Set(row, col, product)
		}
	}

	return &newM

}
