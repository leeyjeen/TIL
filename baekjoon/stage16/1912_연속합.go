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

	var sequence = make([]int, n)
	for i := range sequence {
		fmt.Fscanf(reader, "%d ", &sequence[i])
	}

	fmt.Fprintln(writer, solve(sequence))
}

func solve(sequence []int) int {
	var maxVal = make([]int, len(sequence))
	for i := 0; i < len(maxVal); i++ {
		if i == 0 {
			maxVal[i] = sequence[i]
		} else if i == 1 {
			maxVal[i] = int(math.Max(float64(sequence[i]), float64(sequence[i]+sequence[i-1])))
		} else {
			maxVal[i] = int(math.Max(math.Max(float64(maxVal[i-1]+sequence[i]), float64(sequence[i-1]+sequence[i])), float64(sequence[i])))
		}
	}
	var result int = -1000
	for _, v := range maxVal {
		if v > result {
			result = v
		}
	}
	return result
}
