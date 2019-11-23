package main

//Matrix ...
type Matrix [][]float64

//IdentityMatrix ...
var IdentityMatrix = Matrix(
	[][]float64{
		[]float64{1, 0, 0, 0},
		[]float64{0, 1, 0, 0},
		[]float64{0, 0, 1, 0},
		[]float64{0, 0, 0, 1},
	},
)

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

//Row returns the row of the Matrix
func (matrix *Matrix) Row(r int) []float64 {
	return (*matrix)[r]
}

//Col returns the col of the matrix
func (matrix *Matrix) Col(c int) []float64 {
	h, _ := matrix.Size()
	col := make([]float64, h, h)
	for i, row := range *matrix {
		col[i] = row[c]
	}
	return col
}

//MulMatrix multiplies two 4x4 matrices together
func (matrix *Matrix) MulMatrix(other *Matrix) *Matrix {
	newM := NewMatrix(4, 4)

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			product := zipSum(matrix.Row(row), other.Col(col))
			newM.Set(row, col, product)
		}
	}

	return &newM

}

//MulTuple multiplies a Matrix by a Tuple
func (matrix *Matrix) MulTuple(tuple *Tuple) *Tuple {
	vals := []float64{tuple.x, tuple.y, tuple.z, tuple.w}
	newTup := &Tuple{
		zipSum(matrix.Row(0), vals),
		zipSum(matrix.Row(1), vals),
		zipSum(matrix.Row(2), vals),
		zipSum(matrix.Row(3), vals),
	}

	return newTup

}

//Transpose transposes a matrix (rows become cols and cols become rows)
func (matrix *Matrix) Transpose() *Matrix {
	height, width := matrix.Size()
	newM := NewMatrix(width, height)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			newM.Set(i, j, matrix.Get(j, i))
		}
	}
	return &newM
}
