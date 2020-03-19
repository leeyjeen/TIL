package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n)
	defer writer.Flush()

	fmt.Fprintln(writer, int(math.Pow(2, float64(n)))-1)
	hanoi(n, 1, 3, 2, writer)
}

func hanoi(n, from, to, via int, writer *bufio.Writer) {
	if n == 1 {
		fmt.Fprintf(writer, "%d %d\n", from, to)
		return
	}
	hanoi(n-1, from, via, to, writer)
	move(from, to, writer)
	hanoi(n-1, via, to, from, writer)
}

func move(from, to int, writer *bufio.Writer) {
	fmt.Fprintf(writer, "%d %d\n", from, to)
}
