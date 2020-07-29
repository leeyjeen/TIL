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
	var operCounts = make([]int, n+1)

	for i := 2; i < n+1; i++ {
		if i == 2 || i == 3 {
			operCounts[i] = 1
			continue
		}
		operCounts[i] = operCounts[i-1] + 1
		if i%3 == 0 {
			operCounts[i] = int(math.Min(float64(operCounts[i/3]+1), float64(operCounts[i])))
		}
		if i%2 == 0 {
			operCounts[i] = int(math.Min(float64(operCounts[i/2]+1), float64(operCounts[i])))
		}
	}
	fmt.Fprintln(writer, operCounts[n])
}
