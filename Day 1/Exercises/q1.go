package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows     int
	columns  int
	elements [][]int
}

func (m Matrix) displayRows() string {
	return fmt.Sprintf("Number of Rows: %d\n", m.rows)
}

func (m Matrix) getRows() int {
	return m.rows
}

func (m Matrix) displayColumns() string {
	return fmt.Sprintf("Number of Columns: %d\n", m.columns)
}

func (m Matrix) getColumns() int {
	return m.columns
}

func (m *Matrix) initMatrix() {

	m.elements = make([][]int, m.rows) // Create the outer slice
	for i := 0; i < m.rows; i++ {
		m.elements[i] = make([]int, m.columns) // Create inner slices for each row
	}
}

func (m *Matrix) setElement(row, column, value int) error {
	if row < 0 || row >= m.rows || column < 0 || column >= m.columns {
		return fmt.Errorf("index out of Bounds (row: %d, column: %d)", row, column)
	}

	m.elements[row][column] = value
	return nil
}

func (m *Matrix) add(other Matrix) (*Matrix, error) {
	if m.rows != other.rows || m.columns != other.columns {
		return nil, fmt.Errorf("Matrices must have same dimenions")
	}
	result := Matrix{
		rows:    m.rows,
		columns: m.columns,
	}

	result.initMatrix()

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			// result.setElement(i, j, (m.elements[i][j] + other.elements[i][j]))
			result.elements[i][j] = m.elements[i][j] + other.elements[i][j]
		}

	}

	return &result, nil
}

func (m Matrix) toJSON() (string, error) {
	// Convert the matrix's elements to JSON format
	jsonData, err := json.Marshal(m)
	if err != nil {
		return "", fmt.Errorf("error converting matrix to JSON: %v", err)
	}
	return string(jsonData), nil
}

func main() {
	fmt.Println("Matrix!")
	var matrix = Matrix{
		rows:    2,
		columns: 3,
	}
	fmt.Printf("Rows: %d\n", matrix.getRows())
	fmt.Printf("Columns: %d\n", matrix.getColumns())
	fmt.Printf(matrix.displayRows())
	fmt.Printf(matrix.displayColumns())

	matrix.initMatrix()
	matrix.setElement(0, 0, 1)
	matrix.setElement(0, 1, 2)
	matrix.setElement(0, 2, 3)

	fmt.Printf("Matrix after initialization:\n%v\n", matrix.elements)

	err := matrix.setElement(1, 2, 7)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Matrix after initialization:\n%v\n", matrix.elements)

	matrix1 := Matrix{
		rows:    2,
		columns: 3,
	}
	matrix1.initMatrix()
	matrix1.setElement(0, 0, 1)
	matrix1.setElement(0, 1, 2)
	matrix1.setElement(0, 2, 3)
	matrix1.setElement(1, 0, 4)
	matrix1.setElement(1, 1, 5)
	matrix1.setElement(1, 2, 6)

	matrix2 := Matrix{
		rows:    2,
		columns: 3,
	}
	matrix2.initMatrix()
	matrix2.setElement(0, 0, 7)
	matrix2.setElement(0, 1, 8)
	matrix2.setElement(0, 2, 9)
	matrix2.setElement(1, 0, 10)
	matrix2.setElement(1, 1, 11)
	matrix2.setElement(1, 2, 12)

	// Add the two matrices
	sum, err := matrix1.add(matrix2)
	if err != nil {
		fmt.Println("Error adding matrices:", err)
	} else {
		fmt.Printf("Resultant Matrix (Sum):\n%v\n", sum.elements)
	}

	jsonString, err := matrix1.toJSON()
	if err != nil {
		fmt.Println("Error converting matrix to JSON:", err)
	} else {
		fmt.Println("Matrix as JSON:")
		fmt.Println(jsonString)
	}

}
