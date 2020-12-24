// https://www.acmicpc.net/problem/2630
// 쿼드트리를 만드는 문제
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	paper [][]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	paper = make([][]int, n)
	for i := range paper {
		paper[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &paper[i][j])
		}
	}

	blue, white := checkPart(0, 0, n, 0, 0)
	fmt.Fprintln(writer, white)
	fmt.Fprintln(writer, blue)
}

func checkPart(row, col, n, blue, white int) (int, int) {
	first := paper[row][col]
	var isPaper bool = true
	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			if i == row && j == col {
				continue
			}
			if paper[i][j] != first {
				isPaper = false
				break
			}
		}
		if !isPaper {
			break
		}
	}

	if isPaper {
		if first == 1 {
			blue++
		} else {
			white++
		}
		return blue, white
	}

	interval := n / 2
	blue, white = checkPart(row, col, interval, blue, white)
	blue, white = checkPart(row+interval, col, interval, blue, white)
	blue, white = checkPart(row, col+interval, interval, blue, white)
	blue, white = checkPart(row+interval, col+interval, interval, blue, white)
	return blue, white
}
