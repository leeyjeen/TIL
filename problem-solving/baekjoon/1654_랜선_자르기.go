// https://www.acmicpc.net/problem/1654
// 흔히 parametric search라고도 부르는, 이분 탐색을 응용하여 최솟값이나 최댓값을 찾는 테크닉을 배우는 문제
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var k, n int
	fmt.Fscanln(reader, &k, &n)

	var lengths = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscanln(reader, &lengths[i])
	}

	sort.Ints(lengths) // 주어진 랜선 길이중 가장 긴 랜선의 길이를 탐색 범위로 잡기 위해 정렬한다.
	fmt.Fprintln(writer, getMaxLength(n, lengths))
}

func getMaxLength(n int, lengths []int) (maxLength int) {
	start := 1
	end := lengths[len(lengths)-1] // 탐색 범위를 start, end로 지정
	var mid int
	for start <= end {
		mid = (start + end) / 2 // 이분 탐색과 유사한 방식으로 범위를 좁혀 나간다.
		var totalCount int
		for i := 0; i < len(lengths); i++ {
			totalCount += lengths[i] / mid
		}
		if totalCount >= n {
			if maxLength < mid {
				maxLength = mid
				start = mid + 1
			}
		} else if totalCount < n {
			end = mid - 1
		}
	}
	return
}
