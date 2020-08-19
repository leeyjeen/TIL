package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &x)

	var sequence = make([]int, n)
	for i := range sequence {
		fmt.Fscanf(reader, "%d ", &sequence[i])
		if sequence[i] < x {
			fmt.Fprintf(writer, "%d ", sequence[i])
		}
	}
	fmt.Fprint(writer, "\n")
	writer.Flush()
}
