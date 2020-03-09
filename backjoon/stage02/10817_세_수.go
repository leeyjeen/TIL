package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := make([]int, 3)
	fmt.Scanf("%d %d %d", &numbers[0], &numbers[1], &numbers[2])
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	fmt.Println(numbers[1])
}