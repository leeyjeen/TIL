package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n, &m)
	var coins = make([]int, m)
	for i := range(coins) {
		fmt.Fscanf(reader, "%d ", &coins[i])
	}

	fmt.Fprintln(writer, getWays(n, coins))
}

func getWays(n int, coins []int) int {
	var ways = make([]int, n+1)
	ways[0] = 1
	for _, coin := range(coins) {
		for i := 1; i<n+1; i++ {
			if coin <= i {
				ways[i] += ways[i-coin]
			}
		}
	}
	return ways[n]
}