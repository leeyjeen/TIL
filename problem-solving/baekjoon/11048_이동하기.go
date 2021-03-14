// https://www.acmicpc.net/problem/11048
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var candy = make([][]int, n+1)
	for i := 1; i < n+1; i++ {
		candy[i] = make([]int, m+1)
		for j := 1; j < m+1; j++ {
			fmt.Fscanf(reader, "%d ", &candy[i][j])
		}
	}
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1]))) + candy[i][j]
		}
	}
	fmt.Fprintln(writer, dp[n][m])
}
