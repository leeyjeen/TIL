// https://www.acmicpc.net/problem/2644
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	graph  [][]int
	writer *bufio.Writer
	answer int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, from, to, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &from, &to)
	fmt.Fscanln(reader, &m)
	graph = make([][]int, n+1)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		var parent, child int
		fmt.Fscanln(reader, &parent, &child)
		graph[parent][child] = 1
		graph[child][parent] = 1
	}

	dfs(from, to, 0, 0)
	if answer == 0 {
		answer = -1
	}
	fmt.Fprintln(writer, answer)
}

func dfs(from, to, prev, count int) {
	if from == to {
		answer = count
		return
	}
	for i := 0; i < len(graph); i++ {
		if graph[from][i] == 1 && i != prev { // 연결되어 있고 이전 방문한 점이 아닌 경우
			dfs(i, to, from, count+1) // 재귀함수 호출
		}
	}
	return
}
