package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var x = make([]int, n)
	var y = make([]int, n)
	var order = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x[i], &y[i])
	}

	for i := 0; i < n; i++ {
		order[i]++
		for j := 0; j < n; j++ {
			if x[i] < x[j] && y[i] < y[j] {
				order[i]++
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%d ", order[i])
	}
	fmt.Fprintf(writer, "\n")
}
