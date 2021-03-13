// https://www.acmicpc.net/problem/11057
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
	fmt.Fprintln(writer, getAscCount(n))
}

func getAscCount(length int) int {
	var count int
	var dp = make([][]int, length+1) // dp[i][j]: i번째 자리수에 j가 올 때의 오르막 수 경우의 수
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 10)
	}
	for i := 0; i < 10; i++ {
		dp[1][i] = 1 // 1번째 자리수에 i가 올 때의 오르막 수 경우의 수는 각각 1이다
	}
	for i := 1; i < length+1; i++ { // dp[i][j] = dp[i-1][0] + dp[i-1][1] + ... + dp[i-1][j]
		for j := 0; j < 10; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] += dp[i-1][k]
				dp[i][j] %= 10007
			}
		}
	}
	for i := 0; i < 10; i++ {
		count += dp[length][i]
	}
	return count % 10007
}
