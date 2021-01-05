// https://www.acmicpc.net/problem/9095
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

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(reader, &n)

		count := make([]int, n+1)

		if n+1 > 1 {
			count[1] = 1
		}
		if n+1 > 2 {
			count[2] = 2
		}
		if n+1 > 3 {
			count[3] = 4
		}

		for i := 4; i < n+1; i++ {
			count[i] = count[i-1] + count[i-2] + count[i-3]
		}
		fmt.Fprintln(writer, count[n])
	}
}
