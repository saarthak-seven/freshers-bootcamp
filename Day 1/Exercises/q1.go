package main

import (
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

	fmt.Printf("Matrix after initialization:\n%v\n", matrix.elements)

	err := matrix.setElement(1, 2, 7)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Printf("Matrix after initialization:\n%v\n", matrix.elements)

}
