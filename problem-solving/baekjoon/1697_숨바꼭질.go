// https://www.acmicpc.net/problem/1697
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	var visited = make([]int, 100001)
	bfs(visited, n, k)
	fmt.Fprintln(writer, visited[k])
}

func bfs(visited []int, start, end int) {
	var queue = []int{start}
	for len(queue) > 0 {
		curPos := queue[0]
		if curPos == end {
			break
		}
		queue = queue[1:]
		if curPos-1 >= 0 && visited[curPos-1] == 0 {
			queue = append(queue, curPos-1)
			visited[curPos-1] += (visited[curPos] + 1)
		}
		if curPos+1 <= 100000 && visited[curPos+1] == 0 {
			queue = append(queue, curPos+1)
			visited[curPos+1] += (visited[curPos] + 1)
		}
		if curPos*2 <= 100000 && visited[curPos*2] == 0 {
			queue = append(queue, curPos*2)
			visited[curPos*2] += (visited[curPos] + 1)
		}
	}
}
