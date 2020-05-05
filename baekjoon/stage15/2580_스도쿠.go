package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var slice = make([][]int, 9)
	for index := range slice {
		slice[index] = make([]int, 9)
	}

	for i := range slice {
		for j := range slice[i] {
			fmt.Fscanf(reader, "%d ", &slice[i][j])
		}
	}

	if sudoku(slice) {
		for i := range slice {
			for j := range slice[i] {
				fmt.Fprintf(writer, "%d ", slice[i][j])
			}
			fmt.Fprintln(writer)
		}
	}
}

func sudoku(slice [][]int) bool {
	var row, col int

	if !getUnassignedPosition(slice, &row, &col) {
		return true
	}
	for i := 0; i < 9; i++ {
		if isSafe(slice, row, col, i+1) {
			slice[row][col] = i + 1
			if sudoku(slice) {
				return true
			}
			slice[row][col] = 0
		}
	}
	return false
}

func getUnassignedPosition(slice [][]int, rowptr, colptr *int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if slice[i][j] == 0 {
				*rowptr = i
				*colptr = j
				return true
			}
		}
	}
	return false
}

func isSafe(slice [][]int, row, col, num int) bool {
	if !usedInCol(slice, col, num) &&
		!usedInRow(slice, row, num) &&
		!usedInSubGrid(slice, row-row%3, col-col%3, num) {
		return true
	}
	return false
}

func usedInCol(slice [][]int, col, num int) bool {
	for i := 0; i < 9; i++ {
		if slice[i][col] == num {
			return true
		}
	}
	return false
}

func usedInRow(slice [][]int, row, num int) bool {
	for i := 0; i < 9; i++ {
		if slice[row][i] == num {
			return true
		}
	}
	return false
}

func usedInSubGrid(slice [][]int, startRow, startCol, num int) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if slice[i+startRow][j+startCol] == num {
				return true
			}
		}
	}

	return false
}
