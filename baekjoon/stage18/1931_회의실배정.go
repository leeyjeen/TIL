package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscanln(reader, &n)
	var times = make([][]int, n)
	for i := 0; i < n; i++ {
		var start, end int
		fmt.Fscanln(reader, &start, &end)
		times[i] = []int{start, end}
	}
	maxCount := greedyAlgorithm(times, n)
	fmt.Println(maxCount)
}

func greedyAlgorithm(times [][]int, n int) (maxCount int) {
	sort.Slice(times, func(i, j int) bool {
		if times[j][1] > times[i][1] {
			return true
		} else if times[j][1] == times[i][1] {
			return times[j][0] > times[i][0]
		}
		return false
	})
	previousTime := []int{}
	for i := 0; i < n; i++ {
		if i == 0 {
			maxCount++
			previousTime = times[i]
			continue
		}
		if times[i][0] >= previousTime[1] {
			maxCount++
			previousTime = times[i]
		}
	}
	return
}
