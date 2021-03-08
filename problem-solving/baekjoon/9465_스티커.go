// https://www.acmicpc.net/problem/9465
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

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		var scores = make([][]int, 2)
		var dp = make([][]int, 2)
		for j := 0; j < 2; j++ {
			scores[j] = make([]int, n)
			dp[j] = make([]int, n)
			for k := 0; k < n; k++ {
				fmt.Fscanf(reader, "%d ", &scores[j][k])
			}
		}
		dp[0][0] = scores[0][0]
		dp[1][0] = scores[1][0]
		for j := 1; j < n; j++ {
			if j == 1 {
				dp[0][j] = dp[1][j-1] + scores[0][j]
				dp[1][j] = dp[0][j-1] + scores[1][j]
			} else {
				dp[0][j] = int(math.Max(float64(dp[1][j-1]), float64(dp[1][j-2]))) + scores[0][j]
				dp[1][j] = int(math.Max(float64(dp[0][j-1]), float64(dp[0][j-2]))) + scores[1][j]
			}
		}
		fmt.Fprintln(writer, int(math.Max(float64(dp[0][n-1]), float64(dp[1][n-1]))))
	}
}
