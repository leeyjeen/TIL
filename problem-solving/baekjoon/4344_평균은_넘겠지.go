package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var c int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &c)
	defer writer.Flush()

	for i := 0; i < c; i++ {
		var n int
		fmt.Fscanf(reader, "%d ", &n)

		var scores = make([]float64, n)
		var sum, avg float64
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%f ", &scores[j])
			sum += scores[j]
		}
		avg = sum / float64(n)

		var proportion float64
		for _, val := range scores {
			if val > avg {
				proportion += (1 / float64(n))
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", math.Round(proportion*100000)/1000, "%")
	}
}
