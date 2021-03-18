// https://www.acmicpc.net/problem/11722
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
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &sequence[i])
	}
	fmt.Fprintln(writer, longestDecreasingSequenceCount(sequence))
}

func longestDecreasingSequenceCount(sequence []int) (longestCount int) {
	dp := make([]int, len(sequence))          // 각 요소로부터 최장 감소 부분 수열의 길이를 저장한다
	for i := len(sequence) - 1; i >= 0; i-- { // 비교대상을 수열의 마지막 요소부터 정하여 계산한다
		dp[i] = 1
		length := 0
		for j := i + 1; j < len(sequence); j++ {
			if sequence[i] > sequence[j] { // i, j 요소가 감소하는 부분 수열에 해당할 경우
				if length < dp[j] {
					length = dp[j] // j 요소의 최장 감소 부분 수열의 길이가 더 큰 경우 length값을 갱신해준다
				}
			}
		}
		dp[i] += length // i 요소의 최장 감소 부분 수열의 길이에 length를 더한다
		if longestCount < dp[i] {
			longestCount = dp[i]
		}
	}

	return longestCount
}
