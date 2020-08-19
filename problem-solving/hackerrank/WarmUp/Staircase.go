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

	staircase(n, writer)
}

func staircase(n int, writer *bufio.Writer) {
	for i := 1; i < n+1; i++ {
		for j := n - i; j > 0; j-- {
			fmt.Fprint(writer, " ")
		}
		for j := 0; j < i; j++ {
			fmt.Fprint(writer, "#")
		}
		fmt.Fprintln(writer, "")
	}
}
