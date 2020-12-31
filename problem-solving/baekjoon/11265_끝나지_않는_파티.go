// https://www.acmicpc.net/problem/11265
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}

	graph = floydWarshall(graph)

	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)
		if graph[a-1][b-1] <= c {
			fmt.Fprintln(writer, "Enjoy other party")
		} else {
			fmt.Fprintln(writer, "Stay here")
		}
	}
}

func floydWarshall(graph [][]int) [][]int {
	for i := 0; i < len(graph); i++ { // j->k 경로보다 j->i->k 경로가 빠른 경우 갱신
		for j := 0; j < len(graph); j++ {
			for k := 0; k < len(graph); k++ {
				if graph[j][i]+graph[i][k] < graph[j][k] {
					graph[j][k] = graph[j][i] + graph[i][k]
				}
			}
		}
	}
	return graph
}
