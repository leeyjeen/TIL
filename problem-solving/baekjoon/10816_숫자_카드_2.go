// https://www.acmicpc.net/problem/10816
// 이분 탐색으로 값의 개수를 찾아 봅시다.
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

	var n, m int
	fmt.Fscanln(reader, &n)
	var cards = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &cards[i])
	}

	sort.Ints(cards) // 정렬

	fmt.Fscanln(reader, &m)
	var targets = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &targets[i])
		fmt.Fprintf(writer, "%d ", upperBound(cards, targets[i])-lowerBound(cards, targets[i])+1)
	}
	fmt.Fprintln(writer, "")
}

// lowerBound: return the first occurrence of more than the desired value
func lowerBound(array []int, target int) int {
	var startIdx, endIdx, midIdx int
	endIdx = len(array) - 1

	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2
		if array[midIdx] >= target {
			endIdx = midIdx
		} else {
			startIdx = midIdx + 1
		}
	}
	return endIdx
}

// upperBound: return the position where the value larger than the value
// you are looking for appears first.
func upperBound(array []int, target int) int {
	var startIdx, endIdx, midIdx int
	endIdx = len(array) - 1

	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2
		if array[midIdx] > target {
			endIdx = midIdx
		} else {
			startIdx = midIdx + 1
		}
	}

	if array[endIdx] != target {
		endIdx--
	}
	return endIdx
}

/*
// 해시 맵을 이용한 풀이
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
	fmt.Fscanln(reader, &n)
	var cards = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &cards[i])
	}

	cardMap := map[int]int{} // 맵 이용하여 풀기
	for i := 0; i < len(cards); i++ {
		key := cards[i]
		cardMap[key]++
	}

	fmt.Fscanln(reader, &m)
	var targets = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &targets[i])
		fmt.Fprintf(writer, "%d ", cardMap[targets[i]])
	}
	fmt.Fprintln(writer, "")
}
*/
