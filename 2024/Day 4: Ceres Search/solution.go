package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	searchMatrix, err := readList()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	xmasCounter := 0

	// Traverse the matrix
	for row, search := range searchMatrix {
		for col := range search {

			// Horizontal (right)
			if col+3 < len(search) &&
				search[col] == 'X' &&
				search[col+1] == 'M' &&
				search[col+2] == 'A' &&
				search[col+3] == 'S' {
				xmasCounter++
			}

			// Horizontal (left)
			if col >= 3 &&
				search[col] == 'X' &&
				search[col-1] == 'M' &&
				search[col-2] == 'A' &&
				search[col-3] == 'S' {
				xmasCounter++
			}

			// Vertical (down)
			if row+3 < len(searchMatrix) &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row+1][col] == 'M' &&
				searchMatrix[row+2][col] == 'A' &&
				searchMatrix[row+3][col] == 'S' {
				xmasCounter++
			}

			// Vertical (up)
			if row >= 3 &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row-1][col] == 'M' &&
				searchMatrix[row-2][col] == 'A' &&
				searchMatrix[row-3][col] == 'S' {
				xmasCounter++
			}

			// Diagonal (down-right)
			if row+3 < len(searchMatrix) && col+3 < len(search) &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row+1][col+1] == 'M' &&
				searchMatrix[row+2][col+2] == 'A' &&
				searchMatrix[row+3][col+3] == 'S' {
				xmasCounter++
			}

			// Diagonal (up-left)
			if row >= 3 && col >= 3 &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row-1][col-1] == 'M' &&
				searchMatrix[row-2][col-2] == 'A' &&
				searchMatrix[row-3][col-3] == 'S' {
				xmasCounter++
			}

			// Diagonal (down-left)
			if row+3 < len(searchMatrix) && col >= 3 &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row+1][col-1] == 'M' &&
				searchMatrix[row+2][col-2] == 'A' &&
				searchMatrix[row+3][col-3] == 'S' {
				xmasCounter++
			}

			// Diagonal (up-right)
			if row >= 3 && col+3 < len(search) &&
				searchMatrix[row][col] == 'X' &&
				searchMatrix[row-1][col+1] == 'M' &&
				searchMatrix[row-2][col+2] == 'A' &&
				searchMatrix[row-3][col+3] == 'S' {
				xmasCounter++
			}
		}
	}

	fmt.Println("Number of XMAS found:", xmasCounter)

	xMascounter := 0

	rows := len(searchMatrix)
	cols := len(searchMatrix[0])

	for row := 1; row < rows-1; row++ {
		for col := 1; col < cols-1; col++ {
			if isXMas(searchMatrix, row, col) {
				xMascounter++
			}
		}
	}

	fmt.Println("Number of X-MAS patterns found:", xMascounter)
}

func isXMas(matrix []string, row, col int) bool {
	return isMAS(matrix[row-1][col-1], matrix[row][col], matrix[row+1][col+1]) &&
		isMAS(matrix[row+1][col-1], matrix[row][col], matrix[row-1][col+1])
}

func isMAS(a, b, c byte) bool {
	return (a == 'M' && b == 'A' && c == 'S') || (a == 'S' && b == 'A' && c == 'M')
}

func readList() ([]string, error) {
	var searchMatrix []string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter input (press Ctrl+D or Ctrl+Z to end):")

	for scanner.Scan() {
		searchMatrix = append(searchMatrix, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return searchMatrix, nil
}
