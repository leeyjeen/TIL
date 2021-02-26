// https://www.acmicpc.net/problem/17212
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp    []int
	coins = []int{7, 5, 2, 1}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	dp = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = 100000
	}
	for i := 1; i < n+1; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	fmt.Fprintln(writer, dp[n])
}
