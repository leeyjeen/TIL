// https://www.acmicpc.net/problem/2343
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
	var lessons = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &lessons[i])
	}
	fmt.Fprintln(writer, blueray(lessons, m))
}

func blueray(lessons []int, m int) int {
	var low, high int // low: 레슨 중 가장 긴 시간, high: 레슨 길이의 총 합
	for i := 0; i < len(lessons); i++ {
		high += lessons[i]
		if low < lessons[i] {
			low = lessons[i]
		}
	}

	for low <= high {
		bluerayCount := 0
		sum := 0
		mid := (low + high) / 2 // 블루레이 크기
		for i := 0; i < len(lessons); i++ {
			if sum+lessons[i] > mid { // 블루레이 크기보다 커지는 경우, sum을 0으로 초기화, 블루레이 개수 1증가
				sum = 0
				bluerayCount++
			}
			sum += lessons[i]
		}
		bluerayCount++
		if bluerayCount <= m {
			high = mid - 1 // 블루레이 개수가 m 이하인 경우, 블루레이 크기 탐색의 최대 범위를 감소시킨다
		} else {
			low = mid + 1 // 블루레이 개수가 m보다 큰 경우, 블루레이 크기 탐색의 최소 범위를 증가시킨다
		}
	}
	return low
}
