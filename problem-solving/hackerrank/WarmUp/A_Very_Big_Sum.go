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

	var n, sum int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		var val int
		fmt.Fscanf(reader, "%d ", &val)
		sum += val
	}

	fmt.Fprintln(writer, sum)
}
