// https://www.acmicpc.net/problem/20205
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var bitmap = make([][]int, n)
	for i := 0; i < n; i++ {
		bitmap[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &bitmap[i][j])
		}
	}

	var upsampled = make([][]int, n*k)
	for i := 0; i < n*k; i++ {
		upsampled[i] = make([]int, n*k)
	}
	for i := 0; i < n*k; i++ {
		for j := 0; j < n*k; j++ {
			upsampled[i][j] = bitmap[i/k][j/k]
			fmt.Fprintf(writer, "%d ", upsampled[i][j])
		}
		fmt.Fprintln(writer, "")
	}
}
