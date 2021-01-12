// https://www.acmicpc.net/problem/17626
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
	fmt.Fprintln(writer, getSquareCount(n))
}

func getSquareCount(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i < n+1; i++ {
		minCount := n
		for j := 1; j*j <= i; j++ { // i보다 작은 제곱수들의 dp값 중 최소값을 구한다
			if minCount > dp[i-j*j] {
				minCount = dp[i-j*j]
			}
		}
		dp[i] = minCount + 1 // 최소값과 해당 제곱수를 더하는 경우로 1을 더한 값을 dp[i]에 저장한다
	}
	return dp[n]
}
