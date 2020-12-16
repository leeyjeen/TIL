// https://www.acmicpc.net/problem/2468
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

	var n int
	fmt.Fscanln(reader, &n)

	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}
	visited = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}

	var maxNum int
	for h := 1; h <= 100; h++ {
		var numOfSafety int
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if graph[i][j] >= h && !visited[i][j] {
					numOfSafety++
					dfs(i, j, h)
				}
			}
		}
		if numOfSafety > maxNum {
			maxNum = numOfSafety
		}
		initVisited(visited)
	}

	fmt.Fprintln(writer, maxNum)
}

func dfs(row, col, h int) {
	visited[row][col] = true
	if row+1 < len(graph) && graph[row+1][col] >= h && !visited[row+1][col] {
		dfs(row+1, col, h)
	}
	if row-1 >= 0 && graph[row-1][col] >= h && !visited[row-1][col] {
		dfs(row-1, col, h)
	}
	if col+1 < len(graph) && graph[row][col+1] >= h && !visited[row][col+1] {
		dfs(row, col+1, h)
	}
	if col-1 >= 0 && graph[row][col-1] >= h && !visited[row][col-1] {
		dfs(row, col-1, h)
	}
}

func initVisited(visited [][]bool) {
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			visited[i][j] = false
		}
	}
}
