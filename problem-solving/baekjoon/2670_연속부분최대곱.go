// https://www.acmicpc.net/problem/2670
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

	var realNumbers = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &realNumbers[i])
	}
	for i := 1; i < n; i++ {
		mul := realNumbers[i-1] * realNumbers[i]
		if realNumbers[i] < mul { // 이전의 값과 현재 값을 곱한 값이이 현재 값보다 더 크면 곱한 값으로 갱신
			realNumbers[i] = mul
		}
	}
	var answer float64
	for i := 0; i < n; i++ {
		if answer < realNumbers[i] {
			answer = realNumbers[i]
		}
	}
	fmt.Fprintf(writer, "%.3f\n", answer)
}
