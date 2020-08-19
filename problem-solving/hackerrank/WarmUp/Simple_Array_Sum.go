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

	var sum int
	for i := 0; i < n; i++ {
		var value int
		fmt.Fscanf(reader, "%d ", &value)
		sum += value
	}
	fmt.Fprintln(writer, sum)
}
