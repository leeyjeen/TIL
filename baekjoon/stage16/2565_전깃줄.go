package main

import (
	"bufio"
	"os"
	"fmt"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var poleMap = map[int]int{}
	for i:=0; i<n; i++ {
		var key, val int
		fmt.Fscanln(reader, &key, &val)
		poleMap[key] = val
	}

	var keys = []int{}
	for key := range poleMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	var values = []int{}
	for _, key := range keys {
		values = append(values, poleMap[key])
	}

	// LIS
	var lis = make([]int, n)
	for i := range lis {
		lis[i] = 1
	}
	for i := 0; i < len(values); i++ {
		var maxLis int
		for j := 0; j < i; j++ {
			if values[j] < values[i] && lis[j] > maxLis {
				lis[i] = lis[j] + 1
				maxLis = lis[j]
			}
		}
	}	
	var result int
	for i := range lis {
		if lis[i] > result {
			result = lis[i]
		}
	}
	fmt.Fprintln(writer, n-result)
}