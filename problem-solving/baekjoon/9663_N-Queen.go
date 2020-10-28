package main

import (
	"bufio"
	"fmt"
	"os"
)

var cnt int

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)

	var board = make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}

	_ = solveNQueen(board, n, 0)
	fmt.Println(cnt)
}

func solveNQueen(board [][]int, n, col int) bool {
	if col >= n {
		cnt++
		return true
	}
	for i := 0; i < n; i++ {
		if isSafe(board, n, i, col) {
			board[i][col] = 1
			solveNQueen(board, n, col+1)
			board[i][col] = 0
		}
	}
	return false
}

func isSafe(board [][]int, n, row, col int) bool {
	for i := 0; i < col; i++ {
		if board[row][i] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i < n && j >= 0; i, j = i+1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}
	return true
}
