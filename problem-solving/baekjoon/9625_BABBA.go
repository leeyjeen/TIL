// https://www.acmicpc.net/problem/9625
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

	var k int
	fmt.Fscanln(reader, &k)
	a, b := countAB(k)
	fmt.Fprintf(writer, "%d %d\n", a, b)
}

func countAB(k int) (a, b int) {
	a = 1
	for i := 0; i < k; i++ {
		prevA, prevB := a, b
		a += prevB - prevA
		b += prevA
	}
	return a, b
}
