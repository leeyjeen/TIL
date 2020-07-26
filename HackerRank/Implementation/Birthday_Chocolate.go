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

	var n, d, m int
	fmt.Fscanln(reader, &n)
	var s = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &s[i])
	}
	fmt.Fscanln(reader, &d, &m)

	var count int
	for i := 0; i < n-m+1; i++ {
		var sum int
		for j := i; j < m+i; j++ {
			sum += s[j]
		}
		if sum == d {
			count++
		}
	}
	fmt.Fprintln(writer, count)
}
