// https://www.acmicpc.net/problem/11653
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

	for factor := 2; factor <= n; {
		if n%factor == 0 {
			n /= factor
			fmt.Fprintln(writer, factor)
		} else {
			factor++
		}
	}
}
