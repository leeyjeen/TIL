// https://www.acmicpc.net/problem/9372
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

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)
		for j := 0; j < m; j++ {
			var a, b int
			fmt.Fscanln(reader, &a, &b)
		}
		fmt.Fprintln(writer, n-1) // 스패닝트리: n개의 정점을 가지는 그래프의 최소 간선의 수는 (n-1)개이다.
	}
}
