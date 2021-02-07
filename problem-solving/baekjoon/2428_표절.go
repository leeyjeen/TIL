// https://www.acmicpc.net/problem/2428
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

	var n int
	fmt.Fscanln(reader, &n)

	var sizes = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%f ", &sizes[i])
	}
	sort.Float64s(sizes)
	var totalCount int
	for i := 0; i < len(sizes)-1; i++ {
		totalCount += (getMaxIndex(i, sizes) - i)
	}
	fmt.Fprintln(writer, totalCount)
}

// 검사해야하는 파일의 최대 범위 인덱스 리턴
func getMaxIndex(i int, sizes []float64) int {
	low, high := i+1, len(sizes)-1 // 최초 검사 범위 지정
	val := sizes[i]                // 비교할 값
	mid := 0
	maxIndex := i
	for low <= high {
		mid = (low + high) / 2
		if sizes[mid]*0.9 <= val { // mid값이 조건 만족하는 경우
			maxIndex = mid
			low = mid + 1 // mid값보다 더 높은 범위의 값을 탐색하도록 low 증가
		} else {
			high = mid - 1 // mid값보다 낮은 범위의 값을 탐색하도록 high 감소
		}
	}
	return maxIndex
}
