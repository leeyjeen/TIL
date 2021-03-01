// https://www.acmicpc.net/problem/14606
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp []int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	dp = make([]int, n+1)
	fmt.Fprintln(writer, joy(n))
}

func joy(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	dp[2] = 1
	dp[n] = (n/2)*(n-n/2) + joy(n/2) + joy(n-n/2)
	return dp[n]
}
