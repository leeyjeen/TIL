// https://www.acmicpc.net/problem/1037
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

	var realFactorCount int
	fmt.Fscanln(reader, &realFactorCount)

	var minFactor int = 1000000
	var maxFactor int = 2
	for i := 0; i < realFactorCount; i++ {
		var input int
		fmt.Fscanf(reader, "%d", &input)
		if minFactor > input {
			minFactor = input
		}
		if maxFactor < input {
			maxFactor = input
		}
	}
	fmt.Fprintln(writer, minFactor*maxFactor)
}
