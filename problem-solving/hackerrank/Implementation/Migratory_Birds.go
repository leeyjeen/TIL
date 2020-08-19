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
	var arr = make([]int, n)
	var counts = make([]int, 6)
	var max, mostCommon int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
		counts[arr[i]]++
	}
	for i, v := range counts {
		if v > max {
			max = v
			mostCommon = i
		}
	}
	fmt.Fprintln(writer, mostCommon)
}
