// https://www.acmicpc.net/problem/10708
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

	var n, m int
	fmt.Fscanln(reader, &n)
	fmt.Fscanln(reader, &m)

	var targets = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &targets[i])
	}

	var scores = make([]int, n)
	for i := 0; i < m; i++ {
		var target = targets[i]
		var writes = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &writes[j])
			if writes[j] == target {
				scores[j]++
			} else {
				scores[target-1]++
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, scores[i])
	}
}
