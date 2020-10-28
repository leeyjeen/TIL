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
	var reverseLis = make([]int, n)
	for i := 0; i < n; i++ {
		lis[i] = 1
		reverseLis[i] = 1
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
	for i := len(sequence) - 1; i >= 0; i-- {
		var maxLis int
		for j := len(sequence) - 1; j > i; j-- {
			if sequence[j] < sequence[i] && reverseLis[j] > maxLis {
				reverseLis[i] = reverseLis[j] + 1
				maxLis = reverseLis[j]
			}
		}
	}
	var result int
	for i := 0; i < n; i++ {
		var sum = lis[i] + reverseLis[i] - 1
		if sum > result {
			result = sum
		}
	}
	fmt.Fprintln(writer, result)
}
