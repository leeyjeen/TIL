package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numbers = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &numbers[i])
	}

	var sortedNumbers = make([]int, n)
	mergeSort(numbers, sortedNumbers, 0, n-1) // 병합정렬
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, numbers[i])
	}
}

func merge(numbers, sortedNumbers []int, left, mid, right int) {
	var i = left
	var j = mid + 1
	var k = left

	for true {
		if !(i <= mid && j <= right) {
			break
		}
		if numbers[i] <= numbers[j] {
			sortedNumbers[k] = numbers[i]
			i++
		} else {
			sortedNumbers[k] = numbers[j]
			j++
		}
		k++
	}

	if i > mid {
		for l := j; l <= right; l++ {
			sortedNumbers[k] = numbers[l]
			k++
		}
	} else {
		for l := i; l <= mid; l++ {
			sortedNumbers[k] = numbers[l]
			k++
		}
	}

	for l := left; l <= right; l++ {
		numbers[l] = sortedNumbers[l]
	}
}

func mergeSort(numbers, sortedNumbers []int, left, right int) {
	if left < right {
		var mid = (left + right) / 2
		mergeSort(numbers, sortedNumbers, left, mid)
		mergeSort(numbers, sortedNumbers, mid+1, right)
		merge(numbers, sortedNumbers, left, mid, right)
	}
}
