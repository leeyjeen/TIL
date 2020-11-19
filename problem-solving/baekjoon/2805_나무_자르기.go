// https://www.acmicpc.net/problem/2805
// 이분 탐색을 응용하여 최솟값이나 최댓값을 찾는 문제 2
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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var trees = make([]int, n)
	var maxTree int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &trees[i])
		if maxTree < trees[i] {
			maxTree = trees[i]
		}
	}

	fmt.Fprintln(writer, getMaxHeight(m, maxTree, trees))
}

func getMaxHeight(m, maxTree int, trees []int) (maxHeight int) {
	start := 1
	end := maxTree // 탐색 범위 설정
	var mid int

	for start <= end {
		mid = (start + end) / 2 // 이분 탐색과 유사한 방식으로 범위를 좁혀 나간다.
		totalHeight := 0
		for i := 0; i < len(trees); i++ {
			if trees[i] > mid {
				totalHeight += (trees[i] - mid)
			}
		}
		if totalHeight >= m {
			if maxHeight < mid {
				maxHeight = mid
				start = mid + 1
			}
		} else {
			end = mid - 1
		}
	}
	return
}
