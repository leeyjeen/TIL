// https://www.acmicpc.net/problem/2294
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)
	var coins = make([]int, n)
	var dp = make([]int, k+1)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &coins[i])
	}
	for i := 1; i < k+1; i++ {
		dp[i] = 100000
	}
	for i := 1; i < len(dp); i++ {
		for j := 0; j < n; j++ {
			coin := coins[j]
			if i-coin >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}
	if dp[k] == 100000 {
		fmt.Fprintln(writer, -1)
	} else {
		fmt.Fprintln(writer, dp[k])
	}
}
