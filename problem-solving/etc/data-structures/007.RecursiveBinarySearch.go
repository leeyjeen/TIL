package main

import "fmt"

func BinarySearchRecur(arr []int, first, last, target int) int {
	if first > last {
		return -1
	}
	var mid = int((first+last)/2)
	if arr[mid] == target {
		return mid
	} else if target < arr[mid] {
		return BinarySearchRecur(arr, first, mid-1, target)
	} else {
		return BinarySearchRecur(arr, mid+1, last, target)
	}
}

func main() {
	var arr = []int{1,3,5,7,9}
	var idx int
	idx = BinarySearchRecur(arr, 0, len(arr)-1, 7)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}

	idx = BinarySearchRecur(arr, 0, len(arr)-1, 4)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}
}