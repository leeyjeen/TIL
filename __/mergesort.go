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

	arr = mergeSort(arr)
	fmt.Fprintln(writer, arr)
}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	middle := int(length / 2)
	left := make([]int, middle)
	right := make([]int, length-middle)
	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}
	for j := 0; j < len(left); j++ {
		merged[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		merged[i] = right[j]
		i++
	}
	return merged
}
