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
	var amounts = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &amounts[i])
	}

	var maxAmounts = make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			maxAmounts[i] = amounts[i]
		} else if i == 1 {
			maxAmounts[i] = maxAmounts[i-1] + amounts[i]
		} else if i == 2 {
			maxAmounts[i] = int(math.Max(math.Max(float64(amounts[i-2]+amounts[i]), float64(amounts[i-1]+amounts[i])), float64(maxAmounts[i-1])))
		} else {
			maxAmounts[i] = int(math.Max(math.Max(float64(maxAmounts[i-2]+amounts[i]), float64(maxAmounts[i-3]+amounts[i-1]+amounts[i])), float64(maxAmounts[i-1])))
		}
	}
	fmt.Fprintln(writer, maxAmounts[n-1])
}
