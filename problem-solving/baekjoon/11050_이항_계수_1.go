// https://www.acmicpc.net/problem/11050
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

	var n, k float64
	fmt.Fscanln(reader, &n, &k)

	var result = factorial(n) / factorial(n-k) / factorial(k)
	fmt.Fprintln(writer, result)
}

func factorial(num float64) float64 {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
