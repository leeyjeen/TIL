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

	quickSort(arr, 0, n-1)
	fmt.Fprintln(writer, arr)
}

func quickSort(arr []int, start, end int) {
	if start < end {
		pi := partition(arr, start, end)
		quickSort(arr, start, pi-1)
		quickSort(arr, pi+1, end)
	}
}

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	i := start - 1
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[end] = arr[end], arr[i+1]
	return i + 1
}
