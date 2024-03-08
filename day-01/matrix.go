package day01

import (
	"encoding/json"
	"errors"
)

// Matrix struct represents a matrix data structure
type Matrix struct {
	rows int
	cols int
	data [][]float64
}

// NewMatrix creates a new Matrix with the given rows and columns
func NewMatrix(rows, cols int) *Matrix {
	return &Matrix{rows: rows, cols: cols, data: make([][]float64, rows)}
}

// GetRows returns the number of rows in the matrix
func (m *Matrix) GetRows() int {
	return m.rows
}

// GetCols returns the number of columns in the matrix
func (m *Matrix) GetCols() int {
	return m.cols
}

// SetElement sets the element at the given row (i) and column (j) to the specified value
func (m *Matrix) SetElement(i, j int, value float64) error {
	if i < 0 || i >= m.rows || j < 0 || j >= m.cols {
		return errOutOfBounds
	}
	m.data[i][j] = value
	return nil
}

// Add adds the current matrix with another matrix of the same dimensions
func (m *Matrix) Add(other *Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.cols != other.cols {
		return nil, errMatrixDimMismatch
	}
	result := NewMatrix(m.rows, m.cols)
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.cols; j++ {
			result.data[i][j] = m.data[i][j] + other.data[i][j]
		}
	}
	return result, nil
}

// ToJSON returns the matrix data as a JSON string
func (m *Matrix) ToJSON() (string, error) {
	data, err := json.Marshal(m.data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var (
	errOutOfBounds       = errors.New("index out of bounds")
	errMatrixDimMismatch = errors.New("matrices must have the same dimensions for addition")
)

func RunMatrix() {
	// Example usage
	matrix1 := NewMatrix(2, 3)
	matrix1.SetElement(0, 0, 1)
	matrix1.SetElement(0, 1, 2)
	matrix1.SetElement(0, 2, 3)
	matrix1.SetElement(1, 0, 4)
	matrix1.SetElement(1, 1, 5)
	matrix1.SetElement(1, 2, 6)

	matrix2 := NewMatrix(2, 3)
	matrix2.SetElement(0, 0, 10)
	matrix2.SetElement(0, 1, 20)
	matrix2.SetElement(0, 2, 30)
	matrix2.SetElement(1, 0, 40)
	matrix2.SetElement(1, 1, 50)
	matrix2.SetElement(1, 2, 60)

	resultMatrix, err := matrix1.Add(matrix2)
	if err != nil {
		panic(err)
	}

	jsonData, err := resultMatrix.ToJSON()
	if err != nil {
		panic(err)
	}
	println(jsonData)
}
