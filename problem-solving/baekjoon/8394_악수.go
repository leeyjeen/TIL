// https://www.acmicpc.net/problem/8394
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
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % 10 // dp[i-1]: 마지막 사람이 악수를 하지 않는 경우의 수, dp[i-2]: 마지막 사람이 악수를 하는 경우의 수
	}
	fmt.Fprintln(writer, dp[n])
}
