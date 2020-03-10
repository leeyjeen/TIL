package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, n-i)
	}
	writer.Flush()
}
