// https://www.acmicpc.net/problem/2635
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

	var maxCount int
	var answer []int
	for second := 1; second <= n; second++ {
		var numbers = []int{n, second}
		for i := 2; numbers[i-2]-numbers[i-1] >= 0; i++ {
			numbers = append(numbers, numbers[i-2]-numbers[i-1])
		}
		if maxCount < len(numbers) {
			maxCount = len(numbers)
			answer = numbers
		}
	}
	fmt.Fprintln(writer, maxCount)
	for i := 0; i < len(answer); i++ {
		fmt.Fprintf(writer, "%d ", answer[i])
	}
	fmt.Fprintln(writer, "")
}
