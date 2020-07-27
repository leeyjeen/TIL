package main

import "fmt"

func BinarySearch(arr []int, len, target int) int {
	var first, mid, last int
	last = len-1
	
	for first <= last {
		mid = int((first + last) / 2)

		if target == arr[mid] {
			return mid
		}

		if target < arr[mid] {
			last = mid-1
		} else {
			first = mid+1
		}
	}
	return -1
}

func main() {
	var arr = []int{1,3,5,7,9}
	var idx int
	idx = BinarySearch(arr, len(arr), 7)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}

	idx = BinarySearch(arr, len(arr), 4)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}
}