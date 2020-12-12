// https://www.acmicpc.net/problem/4963
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited [][]bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var w, h int
		fmt.Fscanln(reader, &w, &h)

		if w == 0 && h == 0 {
			break
		}

		graph = make([][]int, h)
		for i := 0; i < h; i++ {
			graph[i] = make([]int, w)
		}
		visited = make([][]bool, h)
		for i := 0; i < h; i++ {
			visited[i] = make([]bool, w)
		}
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Fscanf(reader, "%d ", &graph[i][j])
			}
		}

		var numOfIslands int
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if !visited[i][j] && graph[i][j] == 1 { // 섬이면서 방문하지 않은 경우
					numOfIslands++
					dfs(i, j)
				}
			}
		}
		fmt.Fprintln(writer, numOfIslands)
	}
}

func dfs(h, w int) {
	visited[h][w] = true
	var startH, startW, endH, endW = h, w, h, w
	if h-1 >= 0 {
		startH--
	}
	if h+1 < len(graph) {
		endH++
	}
	if w-1 >= 0 {
		startW--
	}
	if w+1 < len(graph[h]) {
		endW++
	}

	// 상, 하, 좌, 우, 대각선 탐색
	for i := startH; i <= endH; i++ {
		for j := startW; j <= endW; j++ {
			if graph[i][j] == 1 && !visited[i][j] { // 섬이면서 방문하지 않은 경우 재귀
				dfs(i, j)
			}
		}
	}
}
