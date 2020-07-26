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
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	arr = countingSort(arr)
	fmt.Fprintln(writer, arr)
}

func countingSort(arr []int) []int {
	length := len(arr)
	sorted := make([]int, length)
	var maxCount = 100
	count := make([]int, maxCount)
	for i := 0; i < length; i++ {
		count[arr[i]]++
	}
	for i := 1; i < maxCount; i++ {
		count[i] += count[i-1]
	}
	for i := length - 1; i >= 0; i-- {
		sorted[count[arr[i]]-1] = arr[i]
		count[arr[i]]--
	}
	return sorted
}
