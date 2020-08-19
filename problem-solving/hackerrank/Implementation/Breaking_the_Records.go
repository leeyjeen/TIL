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
	var scores = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &scores[i])
	}

	var max, min int
	hitMax, hitMin := 0, 0
	for i, v := range scores {
		if i == 0 {
			max, min = v, v
			continue
		}
		if v > max {
			hitMax++
			max = v
		}
		if v < min {
			hitMin++
			min = v
		}
	}
	fmt.Fprintln(writer, hitMax, hitMin)
}
