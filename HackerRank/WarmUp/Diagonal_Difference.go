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

	var arr = make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &arr[i][j])
		}
	}

	var leftSum, rightSum int
	for i := 0; i < n; i++ {
		leftSum += arr[i][i]
		rightSum += arr[i][n-i-1]
	}

	diff := leftSum - rightSum
	result := diff
	if result < 0 {
		result = -result
	}
	fmt.Fprintln(writer, result)
}
