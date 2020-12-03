// https://www.acmicpc.net/problem/2606
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph   [][]int
	visited []bool
	writer  *bufio.Writer
	count   int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, edge int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &edge)

	graph = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, n+1)
	}
	visited = make([]bool, n+1)

	for i := 0; i < edge; i++ {
		var computer1, computer2 int
		fmt.Fscanln(reader, &computer1, &computer2)
		graph[computer1][computer2] = 1
		graph[computer2][computer1] = 1
	}

	dfs(1)
	fmt.Fprintln(writer, count-1)
}

func dfs(start int) {
	visited[start] = true
	count++
	for i := 0; i < len(graph[start]); i++ {
		if graph[start][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}
