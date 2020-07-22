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
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &arr[i])
	}

	var totals [3]float64
	for _, v := range arr {
		val := float64(v)
		if val > 0 {
			totals[0]++
		} else if val < 0 {
			totals[1]++
		} else {
			totals[2]++
		}
	}

	for _, v := range totals {
		fmt.Fprintf(writer, "%f\n", v/float64(n))
	}
}
