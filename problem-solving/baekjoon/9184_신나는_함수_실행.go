// https://www.acmicpc.net/problem/9184
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	dp [][][]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)
		if a == -1 && b == -1 && c == -1 {
			break
		}
		dp = make([][][]int, 21)
		for i := 0; i < 21; i++ {
			dp[i] = make([][]int, 21)
			for j := 0; j < 21; j++ {
				dp[i][j] = make([]int, 21)
			}
		}
		fmt.Fprintf(writer, "w(%d, %d, %d) = %d\n", a, b, c, w(a, b, c))
	}
}

func w(a, b, c int) int {
	if a <= 0 || b <= 0 || c <= 0 {
		return 1
	}
	if a > 20 || b > 20 || c > 20 {
		return w(20, 20, 20)
	}
	if dp[a][b][c] != 0 {
		return dp[a][b][c]
	}
	if a < b && b < c {
		dp[a][b][c] = w(a, b, c-1) + w(a, b-1, c-1) - w(a, b-1, c)
	} else {
		dp[a][b][c] = w(a-1, b, c) + w(a-1, b-1, c) + w(a-1, b, c-1) - w(a-1, b-1, c-1)
	}
	return dp[a][b][c]
}
