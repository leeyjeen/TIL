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

	var sequence = make([]int, n)
	for i := range sequence {
		fmt.Fscanf(reader, "%d ", &sequence[i])
	}

	var lis = make([]int, n)
	for i := range lis {
		lis[i] = 1
	}
	for i := 0; i < len(sequence); i++ {
		var maxLis int
		for j := 0; j < i; j++ {
			if sequence[j] < sequence[i] && lis[j] > maxLis {
				lis[i] = lis[j] + 1
				maxLis = lis[j]
			}
		}
	}

	var result int
	for i := range lis {
		if lis[i] > result {
			result = lis[i]
		}
	}
	fmt.Fprintln(writer, result)
}
