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

	var ar = make([]int, n)
	var max, maxCount int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &ar[i])
		if ar[i] > max {
			max = ar[i]
			maxCount = 1
		} else if ar[i] == max {
			maxCount++
		}
	}
	fmt.Fprintln(writer, maxCount)
}
