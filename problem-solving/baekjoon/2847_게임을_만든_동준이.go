// https://www.acmicpc.net/problem/2847
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
	var scores = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &scores[i])
	}
	var result int
	for i := n - 2; i >= 0; i-- {
		if scores[i+1] <= scores[i] {
			diff := scores[i] - scores[i+1] + 1
			result += diff
			scores[i] -= diff
		}
	}
	fmt.Fprintln(writer, result)
}
