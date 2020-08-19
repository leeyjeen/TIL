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

	var s, t, a, b, m, n int
	fmt.Fscanln(reader, &s, &t)
	fmt.Fscanln(reader, &a, &b)
	fmt.Fscanln(reader, &m, &n)
	var apple, orange, diff int
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &diff)
		if a+diff >= s && a+diff <= t {
			apple++
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &diff)
		if b+diff >= s && b+diff <= t {
			orange++
		}
	}
	fmt.Fprintf(writer, "%d\n%d\n", apple, orange)
}
