// https://www.acmicpc.net/problem/1182
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	count int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscanln(reader, &n, &s)
	var seq = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &seq[i])
	}
	computeSubsequenceCount(0, 0, s, seq)
	fmt.Fprintln(writer, count)
}

func computeSubsequenceCount(index, sum, aim int, seq []int) {
	if index > len(seq)-1 {
		return
	}
	sum += seq[index]
	if sum == aim {
		count++
	}
	computeSubsequenceCount(index+1, sum, aim, seq)            // 현재 인덱스 값을 더하지 않는 경우
	computeSubsequenceCount(index+1, sum-seq[index], aim, seq) // 현재 인덱스 값을 더하는 경우
}
