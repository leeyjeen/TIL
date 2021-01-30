// https://www.acmicpc.net/problem/2217
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

	var ropes = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &ropes[i])
	}
	sort.Ints(ropes)
	var maxWeight int
	for i := len(ropes) - 1; i >= 0; i-- {
		weight := ropes[i] * (len(ropes) - i)
		if weight > maxWeight {
			maxWeight = weight
		}
	}
	fmt.Fprintln(writer, maxWeight)
}
