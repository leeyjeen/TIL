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
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &p[i])
	}
	fmt.Println(greedyAlgorithm(n, p))
}

func greedyAlgorithm(n int, p []int) (minTime int) {
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })
	minTimes := []int{p[0]}
	minTime = p[0]
	for i := 1; i < n; i++ {
		minTimes = append(minTimes, minTimes[i-1]+p[i])
		minTime += minTimes[i]
	}
	return
}
