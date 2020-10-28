package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int

	// user bufio for fast scan
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, i+1)
	}
	writer.Flush()
}
