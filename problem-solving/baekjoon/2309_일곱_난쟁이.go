// https://www.acmicpc.net/problem/2309
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

	heights := []int{}
	var sum int
	for i := 0; i < 9; i++ {
		var h int
		fmt.Fscanln(reader, &h)
		heights = append(heights, h)
		sum += h
	}

	sort.Ints(heights)
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			if heights[i]+heights[j] == sum-100 {
				for k := 0; k < 9; k++ {
					if k != i && k != j {
						fmt.Fprintln(writer, heights[k])
					}
				}
				return
			}
		}
	}
}
