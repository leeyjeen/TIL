// https://www.acmicpc.net/problem/10994
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stars [][]string
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	stars = make([][]string, 4*n-3)
	for i := 0; i < len(stars); i++ {
		stars[i] = make([]string, 4*n-3)
		for j := 0; j < len(stars[i]); j++ {
			stars[i][j] = " "
		}
	}

	drawStar(n, 0, 0)

	for i := 0; i < len(stars); i++ {
		fmt.Fprintln(writer, strings.Join(stars[i], ""))
	}
}

func drawStar(n, row, col int) {
	length := 4*n - 3 // 별을 그릴 사각형의 변의 길이
	if length == 1 {
		stars[row][col] = "*"
		return
	}

	for r := row; r < row+length; r++ { // 세로변 그리기
		stars[r][col] = "*"
		stars[r][col+length-1] = "*"
	}

	for c := col + 1; c < col+length-1; c++ { // 가로변 그리기
		stars[row][c] = "*"
		stars[row+length-1][c] = "*"
	}

	drawStar(n-1, row+2, col+2) // 내부의 사각형 별 그리기
}
