package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var m, n int
		fmt.Fscanln(reader, &m)
		fmt.Fscanln(reader, &n)
		var cost = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &cost[j])
		}
		index1, index2 := search(m, cost)
		fmt.Fprintln(writer, index1, index2)
	}
}

func search(m int, cost []int) (index1, index2 int) {
	for i := 0; i < len(cost)-1; i++ {
		costSunny := cost[i]
		if costSunny >= m {
			continue
		}
		index1 = i + 1
		for j := i + 1; j < len(cost); j++ {
			costJohnny := cost[j]
			if m == costSunny+costJohnny {
				index2 = j + 1
				return
			}
		}
	}
	return
}
