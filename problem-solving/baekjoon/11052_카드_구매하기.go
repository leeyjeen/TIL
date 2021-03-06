// https://www.acmicpc.net/problem/11052
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

	var n int
	fmt.Fscanln(reader, &n)
	var p = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscanf(reader, "%d ", &p[i])
	}
	fmt.Fprintln(writer, getMaxPrice(p))
}

func getMaxPrice(p []int) int {
	n := len(p)
	var dp = make([]int, n)
	dp[1] = p[1]
	for i := 2; i < n; i++ {
		dp[i] = p[i]
		for j := 1; j < i; j++ {
			dp[i] = int(math.Max(float64(dp[i]), float64(dp[i-j]+p[j])))
		}
	}
	return dp[n-1]
}
