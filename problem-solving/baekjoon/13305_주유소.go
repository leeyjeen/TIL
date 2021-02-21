// https://www.acmicpc.net/problem/13305
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

	var n int
	fmt.Fscanln(reader, &n)

	var distances = make([]int, n-1)
	var prices = make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Fscanf(reader, "%d ", &distances[i])
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &prices[i])
	}

	result := 0
	for i := 0; i < n-1; i++ {
		cur := prices[i]
		dist := distances[i]
		for j := i + 1; j < n-1; j++ {
			next := prices[j]
			if cur < next {
				dist += distances[j]
				i++
			} else {
				break
			}
		}
		result += dist * cur
	}
	fmt.Fprintln(writer, result)
}
