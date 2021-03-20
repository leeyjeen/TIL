// https://www.acmicpc.net/problem/11055
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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}

	fmt.Fprintln(writer, longestAscendingSequenceSum(a))
}

func longestAscendingSequenceSum(sequence []int) (maxSum int) {
	dp := make([]int, len(sequence)) // 각 요소까지의 가장 큰 증가 부분 수열의 합을 저장한다
	for i := 0; i < len(sequence); i++ {
		maxDp := 0
		for j := 0; j < i; j++ {
			if sequence[j] < sequence[i] { // j, i 요소가 증가하는 부분 수열에 해당할 경우
				if maxDp < dp[j] {
					maxDp = dp[j] // j 요소까지의 가장 큰 증가 부분 수열의 크기를 maxDp에 갱신한다
				}
			}
		}
		dp[i] = maxDp + sequence[i]
		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}
	return
}
