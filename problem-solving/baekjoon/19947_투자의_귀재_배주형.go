// https://www.acmicpc.net/problem/19947
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

	var h float64
	var y int
	fmt.Fscanln(reader, &h, &y)
	fmt.Fprintln(writer, getMaxProfit(h, y))
}

func getMaxProfit(h float64, y int) float64 {
	var dp = make([]float64, y+1)
	dp[0] = h
	for i := 1; i < y+1; i++ {
		dp[i] = float64(int(dp[i-1] * 1.05))
		if i >= 3 {
			threeYears := float64(int(dp[i-3] * 1.2))
			if threeYears > dp[i] {
				dp[i] = threeYears
			}
		}
		if i >= 5 {
			fiveYears := float64(int(dp[i-5] * 1.35))
			if fiveYears > dp[i] {
				dp[i] = fiveYears
			}
		}
	}
	return dp[y]
}
