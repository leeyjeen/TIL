// https://www.acmicpc.net/problem/11724
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited []bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	graph = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, n+1)
	}
	visited = make([]bool, n+1)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Fscanln(reader, &u, &v)
		graph[u][v] = 1
		graph[v][u] = 1
	}

	var numOfConnected int
	for i := 0; i < n+1; i++ {
		if !visited[i] {
			numOfConnected++
			dfs(i)
		}
	}
	fmt.Fprintln(writer, numOfConnected-1)
}

func dfs(start int) {
	visited[start] = true
	for i := 0; i < len(graph); i++ {
		if !visited[i] && graph[start][i] == 1 {
			dfs(i)
		}
	}
}
