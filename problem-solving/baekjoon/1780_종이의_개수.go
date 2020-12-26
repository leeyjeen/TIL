// https://www.acmicpc.net/problem/1780
// 쿼드트리와 비슷한데 4개 대신 9개로 나누는 문제
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	paper      [][]int
	minusCount int
	zeroCount  int
	plusCount  int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	paper = make([][]int, n)
	for i := 0; i < n; i++ {
		paper[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &paper[i][j])
		}
	}

	checkPaper(0, 0, n)
	fmt.Fprintln(writer, minusCount)
	fmt.Fprintln(writer, zeroCount)
	fmt.Fprintln(writer, plusCount)
}

func checkPaper(row, col, n int) {
	first := paper[row][col]
	isComplete := true
	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			if i == row && j == col {
				continue
			}
			if paper[i][j] != first {
				isComplete = false
				break
			}
		}
		if !isComplete {
			break
		}
	}
	if isComplete {
		switch first {
		case -1:
			minusCount++
		case 0:
			zeroCount++
		case 1:
			plusCount++
		}
		return
	}
	interval := n / 3
	checkPaper(row, col, interval)
	checkPaper(row, col+interval, interval)
	checkPaper(row, col+2*interval, interval)
	checkPaper(row+interval, col, interval)
	checkPaper(row+interval, col+interval, interval)
	checkPaper(row+interval, col+2*interval, interval)
	checkPaper(row+2*interval, col, interval)
	checkPaper(row+2*interval, col+interval, interval)
	checkPaper(row+2*interval, col+2*interval, interval)
}
