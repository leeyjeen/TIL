// https://www.acmicpc.net/problem/2740
// 행렬의 거듭제곱을 계산하기 전에 먼저 풀어야 할 문제
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

	var n, m, k int
	fmt.Fscanln(reader, &n, &m)

	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(reader, "%d ", &a[i][j])
		}
	}

	fmt.Fscanln(reader, &m, &k)
	var b = make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = make([]int, k)
		for j := 0; j < k; j++ {
			fmt.Fscanf(reader, "%d ", &b[i][j])
		}
	}
	var result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, k)
		for j := 0; j < k; j++ {
			for l := 0; l < m; l++ {
				result[i][j] += a[i][l] * b[l][j]
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			fmt.Fprintf(writer, "%d ", result[i][j])
		}
		fmt.Fprintln(writer, "")
	}
}
