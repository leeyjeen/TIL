// https://www.acmicpc.net/problem/20207
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

	var n int
	fmt.Fscanln(reader, &n)

	scheduleMap := map[int]int{}
	for i := 0; i < n; i++ {
		var s, e int
		fmt.Fscanln(reader, &s, &e)
		for j := s; j <= e; j++ {
			scheduleMap[j]++
		}
	}

	keys := []int{}
	for key := range scheduleMap {
		keys = append(keys, key)
	}

	sort.Ints(keys)

	var maxWidth, maxHeight, area int
	for i, val := range keys {
		if i != 0 && keys[i-1]+1 != val {
			area += maxWidth * maxHeight
			maxWidth = 0
			maxHeight = 0
		}
		maxWidth++
		if scheduleMap[val] > maxHeight {
			maxHeight = scheduleMap[val]
		}
	}
	area += maxWidth * maxHeight
	fmt.Fprintln(writer, area)
}
