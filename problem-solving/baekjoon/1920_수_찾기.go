// https://www.acmicpc.net/problem/1920
// 배열을 정렬한 후 이분 탐색으로 값을 찾아 봅시다.
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

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &a[i])
	}

	fmt.Fscanln(reader, &m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &b[i])
	}

	sort.Ints(a) // 정렬
	for i := 0; i < m; i++ {
		if binarySearch(a, b[i]) {
			fmt.Fprintln(writer, 1)
		} else {
			fmt.Fprintln(writer, 0)
		}
	}
}

// binarySearch
// array: sorted array
// val: value to be searched
func binarySearch(array []int, val int) bool {
	low := 0               // lower bound
	high := len(array) - 1 // upper bound

	for low <= high {
		mid := (low + high) / 2

		if array[mid] < val {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if low == len(array) || array[low] != val {
		return false
	}
	return true
}
