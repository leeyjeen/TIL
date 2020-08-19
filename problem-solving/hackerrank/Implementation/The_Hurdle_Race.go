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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var height = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &height[i])
	}
	fmt.Fprintln(writer, hurdleRace(k, height))
}

func hurdleRace(k int, height []int) int {
	var doses, max int
	for _, v := range height {
		if v > max {
			max = v
		}
	}
	if max > k {
		doses = max - k
	}
	return doses
}
