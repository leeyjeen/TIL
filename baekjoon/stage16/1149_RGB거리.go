package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	cost := [][]int{}
	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Fscanln(reader, &a, &b, &c)
		cost = append(cost, []int{a, b, c})
	}
	minVal := calculate(cost, n)
	fmt.Fprintln(writer, minVal)
}

func calculate(cost [][]int, n int) (minVal int) {
	minVal = 0
	if n == 1 {
		minVal = int(math.Min(math.Min(float64(cost[0][0]), float64(cost[0][1])), float64(cost[0][2])))
	}

	minCosts := [][]int{}
	minCosts = append(minCosts, []int{cost[0][0], cost[0][1], cost[0][2]})

	for i := 1; i < n; i++ {
		values := []int{
			int(math.Min(float64(minCosts[i-1][1]), float64(minCosts[i-1][2]))) + cost[i][0],
			int(math.Min(float64(minCosts[i-1][0]), float64(minCosts[i-1][2]))) + cost[i][1],
			int(math.Min(float64(minCosts[i-1][0]), float64(minCosts[i-1][1]))) + cost[i][2],
		}
		minCosts = append(minCosts, values)
	}

	minVal = int(math.Min(math.Min(float64(minCosts[n-1][0]), float64(minCosts[n-1][1])), float64(minCosts[n-1][2])))
	return
}
