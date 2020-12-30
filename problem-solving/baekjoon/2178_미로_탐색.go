// https://www.acmicpc.net/problem/2178
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	graph   [][]string
	count   [][]int
	rowDiff = []int{0, 0, -1, 1}
	colDiff = []int{-1, 1, 0, 0}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	count = make([][]int, n)
	for i := 0; i < n; i++ {
		count[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		input, _ := reader.ReadString('\n') // 공백 포함하여 입력 받기 위해 ReadString() 사용
		inputs := strings.Split(strings.ReplaceAll(input, "\n", ""), "")
		graph = append(graph, inputs)
	}

	bfs(0, 0)
	fmt.Fprintln(writer, count[n-1][m-1])
}

type pos struct {
	row int
	col int
}

func bfs(row, col int) { // 최단 경로 구할 때 BFS 활용
	queue := []pos{pos{row, col}}
	count[row][col] = 1
	for len(queue) > 0 {
		r := queue[0].row
		c := queue[0].col
		if r == len(graph) && c == len(graph[0]) {
			break
		}
		if len(queue) == 1 {
			queue = []pos{}
		} else {
			queue = queue[1:]
		}

		for i := 0; i < 4; i++ {
			newRow := r + rowDiff[i]
			newCol := c + colDiff[i]
			if check(newRow, newCol) {
				queue = append(queue, pos{newRow, newCol})
				count[newRow][newCol] = count[r][c] + 1
			}
		}
	}
}

func check(row, col int) bool {
	if row >= len(graph) || row < 0 || col >= len(graph[0]) || col < 0 {
		return false
	}
	if graph[row][col] == "0" || count[row][col] != 0 {
		return false
	}
	return true
}
