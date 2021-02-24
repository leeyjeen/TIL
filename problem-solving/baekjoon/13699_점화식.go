// https://www.acmicpc.net/problem/13699
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

	var dp = make([]int, n+1)
	dp[0] = 1
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	fmt.Fprintln(writer, dp[n])
}
