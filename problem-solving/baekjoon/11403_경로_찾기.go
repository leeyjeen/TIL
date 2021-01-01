// https://www.acmicpc.net/problem/11403
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

	var n int
	fmt.Fscanln(reader, &n)

	var graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &graph[i][j])
		}
	}

	result := floydWarshall(graph)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result[i]); j++ {
			fmt.Fprintf(writer, "%d ", result[i][j])
		}
		fmt.Fprintln(writer, "")
	}
}

func floydWarshall(graph [][]int) [][]int {
	for i := 0; i < len(graph); i++ { // j->i->k 경로가 존재하는 경우 j->k 경로가 존재하는 것으로 갱신한다.
		for j := 0; j < len(graph); j++ {
			for k := 0; k < len(graph); k++ {
				if graph[j][i] != 0 && graph[i][k] != 0 {
					graph[j][k] = 1
				}
			}
		}
	}
	return graph
}
