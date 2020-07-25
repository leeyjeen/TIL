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

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)
		var b = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &b[j])
		}
		fmt.Fprintln(writer, cost(b))
	}
}

func cost(b []int) int {
	var maxValues = make([][]int, len(b))
	for i := 0; i < len(maxValues); i++ {
		maxValues[i] = make([]int, 2)
	}

	for i := 1; i < len(b); i++ {
		maxValues[i][0] = int(math.Max(float64(maxValues[i-1][0]), float64(maxValues[i-1][1]+b[i-1]-1)))
		maxValues[i][1] = int(math.Max(float64(maxValues[i-1][0]+b[i]-1), float64(maxValues[i-1][1]+int(math.Abs(float64(b[i-1]-b[i]))))))
	}
	return int(math.Max(float64(maxValues[len(b)-1][0]), float64(maxValues[len(b)-1][1])))
}
