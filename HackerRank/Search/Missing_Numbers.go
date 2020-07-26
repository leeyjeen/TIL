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
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}
	fmt.Fscanln(reader, &m)
	var brr = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &brr[i])
	}

	result := missingNumbers(arr, brr)
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer, "")
}

func missingNumbers(arr, brr []int) []int {
	numbersMap := map[int]int{}
	for _, v := range brr {
		if _, ok := numbersMap[v]; ok {
			numbersMap[v]++
		} else {
			numbersMap[v] = 1
		}
	}
	for _, v := range arr {
		numbersMap[v]--
		if numbersMap[v] == 0 {
			delete(numbersMap, v)
		}
	}
	keys := []int{}
	for key := range numbersMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	return keys
}
