// https://www.acmicpc.net/problem/11725
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	tree    [][]int
	visited []bool
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	tree = make([][]int, n+1)
	visited = make([]bool, n+1)
	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Fscanln(reader, &u, &v)
		tree[u] = append(tree[u], v)
		tree[v] = append(tree[v], u)
	}
	parents := bfs()
	for i := 2; i < n+1; i++ {
		fmt.Fprintln(writer, parents[i])
	}
}

func bfs() (parents []int) {
	queue := []int{1}
	parents = make([]int, len(tree))
	for len(queue) > 0 {
		parent := queue[0]
		queue = queue[1:]
		for _, v := range tree[parent] {
			if !visited[v] {
				parents[v] = parent
				queue = append(queue, v)
				visited[v] = true
			}
		}
	}
	return parents
}
