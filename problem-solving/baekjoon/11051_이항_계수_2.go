// https://www.acmicpc.net/problem/11051
// 더 넓은 범위의 이항 계수를 동적 계획법으로 구하는 문제
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	/*
		팩토리얼 공식 대신 파스칼 삼각형을 이용하여 점화식 작성
		nCk = nCn-k
		nCk + nCk+1 = n+1Ck+1
	*/
	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == j || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = (dp[i-1][j] + dp[i-1][j-1]) % 10007
			}
		}
	}

	fmt.Fprintln(writer, dp[n][k])
}
