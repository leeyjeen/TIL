package main

import "fmt"

func LinearSearch(ar []int, len, target int) int {
	for i := 0; i < len; i++ {
		if ar[i] == target {
			return i // return index value of found object
		}
	}
	return -1 // return a value that means it has not been found
}

func main() {
	var arr = []int{3, 5, 2, 4, 9}
	var idx int = LinearSearch(arr, len(arr), 4)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}

	idx = LSearch(arr, len(arr), 7)
	if idx == -1 {
		fmt.Printf("failed to search..\n")
	} else {
		fmt.Printf("target stored index: %d\n", idx)
	}
}
