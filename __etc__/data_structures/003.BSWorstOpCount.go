package main

import "fmt"

func BinarySearch(arr []int, len, target int) int {
	var first = 0
	var last = len - 1
	var mid, opCount int

	for first <= last {
		mid = int((first + last) / 2)

		if target == arr[mid] {
			return mid
		}
		if target < arr[mid] {
			last = mid - 1
		} else {
			first = mid + 1
		}
		opCount++
	}
	fmt.Printf("Number of comparison operations: %d \n", opCount)
	return -1
}

func main() {
	var arr1 = make([]int, 500)
	var arr2 = make([]int, 5000)
	var arr3 = make([]int, 50000)
	var idx int

	idx = BinarySearch(arr1, len(arr1), 1)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}

	idx = BinarySearch(arr2, len(arr2), 2)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}

	idx = BinarySearch(arr3, len(arr3), 3)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}
}
