package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscanln(reader, &n, &k)
	var values = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &values[i])
	}
	minCount := greedyAlgorithm(values, k, n)
	fmt.Println(minCount)
}

func greedyAlgorithm(values []int, k, n int) (coinCount int) {
	for i := n - 1; i >= 0; i-- {
		coinCount = coinCount + k/values[i]
		k %= values[i]
	}
	return
}
