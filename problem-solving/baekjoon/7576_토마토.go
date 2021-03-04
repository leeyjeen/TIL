// https://www.acmicpc.net/problem/7576
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph [][]int
	queue []tomato
	dcol  = []int{-1, 0, 1, 0}
	drow  = []int{0, 1, 0, -1}
)

type tomato struct {
	row int
	col int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var m, n int
	fmt.Fscanln(reader, &m, &n)

	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
			if graph[i][j] == 1 {
				queue = append(queue, tomato{row: i, col: j}) // 익은 토마토의 좌표를 큐에 추가한다.
			}
		}
	}
	bfs()
	fmt.Fprintln(writer, getTimeToGrowRipe(graph))
}

func bfs() {
	for len(queue) > 0 {
		row := queue[0].row
		col := queue[0].col
		queue = queue[1:]

		for i := 0; i < 4; i++ { // 큐에 담긴 순서대로 상하좌우 토마토를 확인한다.
			newRow := row + drow[i]
			newCol := col + dcol[i]
			if newRow < 0 || newRow > len(graph)-1 || newCol < 0 || newCol > len(graph[row])-1 {
				continue
			}
			if graph[newRow][newCol] != 0 {
				continue
			}
			graph[newRow][newCol] += (graph[row][col] + 1) // 익지 않은 토마토 값을 현재 토마토 값 + 1로 갱신한다.
			queue = append(queue, tomato{newRow, newCol})  // 익게 된 토마토 좌표를 큐에 추가한다.
		}
	}
}

func getTimeToGrowRipe(graph [][]int) int {
	var totalTime int
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i]); j++ {
			if graph[i][j] == 0 { // 안 익은 토마토가 있는 경우
				return -1
			}
			if graph[i][j] > totalTime {
				totalTime = graph[i][j]
			}
		}
	}
	return totalTime - 1
}
