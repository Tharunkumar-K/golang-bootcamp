package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int     `json:"rows"`
	Cols     int     `json:"cols"`
	Elements [][]int `json:"elements"`
}

func CreateMatrix(rows, cols int) Matrix {
	matrix := Matrix{
		Rows:     rows,
		Cols:     cols,
		Elements: make([][]int, rows),
	}
	for i := range matrix.Elements {
		matrix.Elements[i] = make([]int, cols)
	}
	return matrix
}

func (m Matrix) GetRows() int {
	return m.Rows
}

func (m Matrix) GetColumns() int {
	return m.Cols
}

func (m Matrix) SetElement(i, j, value int) {
	m.Elements[i][j] = value
}

func (m Matrix) Add(other Matrix) Matrix {
	result := CreateMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

func (m Matrix) ToJSON() (string, error) {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

func main() {
	var rows, cols int

	fmt.Println("Enter the number of rows and columns:")
	_, err := fmt.Scan(&rows, &cols)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	matrix1 := CreateMatrix(rows, cols)
	fmt.Println("Enter elements for Matrix 1:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var value int
			fmt.Printf("Enter element at position (%d,%d): ", i, j)
			_, err := fmt.Scan(&value)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			matrix1.SetElement(i, j, value)
		}
	}

	matrix2 := CreateMatrix(rows, cols)
	fmt.Println("Enter elements for Matrix 2:")
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var value int
			fmt.Printf("Enter element at position (%d,%d): ", i, j)
			_, err := fmt.Scan(&value)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			matrix2.SetElement(i, j, value)
		}
	}

	fmt.Println("Matrix 1:", matrix1)
	fmt.Println("Matrix 1 Rows:", matrix1.GetRows())
	fmt.Println("Matrix 1 Columns:", matrix1.GetColumns())

	fmt.Println("Matrix 2:", matrix2)
	fmt.Println("Matrix 2 Rows:", matrix2.GetRows())
	fmt.Println("Matrix 2 Columns:", matrix2.GetColumns())

	sum := matrix1.Add(matrix2)
	fmt.Println("Sum of matrices:", sum)
	fmt.Println("Matrix 2 Rows:", sum.GetRows())
	fmt.Println("Matrix 2 Columns:", sum.GetColumns())

	matrix1Json, err := matrix1.ToJSON()
	if err != nil {
		fmt.Println("Error converting matrix 1 to JSON:", err)
	} else {
		fmt.Println("Matrix 1 as JSON:", matrix1Json)
	}
	matrix2Json, err := matrix2.ToJSON()
	if err != nil {
		fmt.Println("Error converting matrix 2 to JSON:", err)
	} else {
		fmt.Println("Matrix 2 as JSON:", matrix2Json)
	}
	sumJson, err := sum.ToJSON()
	if err != nil {
		fmt.Println("Error converting matrix sum to JSON:", err)
	} else {
		fmt.Println("Matrix sum as JSON:", sumJson)
	}
}
