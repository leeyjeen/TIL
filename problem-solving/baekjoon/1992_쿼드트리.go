// https://www.acmicpc.net/problem/1992
// 쿼드트리를 문자열로 바꾸는 문제
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	image  [][]string
	writer *bufio.Writer
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		inputs := strings.Split(input, "")
		image = append(image, inputs)
	}

	quadTree(0, 0, n)
	fmt.Fprintln(writer, "")
}

func quadTree(row, col, n int) {
	first := image[row][col]
	var isComplete bool = true

	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			if i == row && j == col {
				continue
			}
			if image[i][j] != first {
				isComplete = false
				break
			}
		}
		if !isComplete {
			break
		}
	}
	if isComplete {
		fmt.Fprint(writer, first)
	} else {
		fmt.Fprint(writer, "(")
		interval := n / 2
		quadTree(row, col, interval)
		quadTree(row, col+interval, interval)
		quadTree(row+interval, col, interval)
		quadTree(row+interval, col+interval, interval)
		fmt.Fprint(writer, ")")
	}
}
