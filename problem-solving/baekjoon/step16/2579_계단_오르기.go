package main

import (
	"bufio"
	"os"
	"fmt"
	"math"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var scores = make([]int, n)
	for i:=0; i<n; i++ {
		fmt.Fscanln(reader, &scores[i])
	}

	var maxValues = make([]int, n)
	for i:=0; i<n; i++ {
		if i == 0 {
			maxValues[i] = scores[i]
		} else if i == 1 {
			maxValues[i] = maxValues[i-1] + scores[i]
		} else if i == 2 {
			maxValues[i] = int(math.Max(float64(scores[i-1] +scores[i]), float64(scores[i-2]+scores[i])))
		} else {
			maxValues[i] = int(math.Max(float64(maxValues[i-2]+scores[i]), float64(maxValues[i-3]+scores[i-1]+scores[i])))
		}
	}
	fmt.Fprintln(writer, maxValues[n-1])
}