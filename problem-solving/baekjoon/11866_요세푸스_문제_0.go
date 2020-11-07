// https://www.acmicpc.net/problem/11866
// 큐를 이용해 제거 과정을 구현하는 문제
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

	var n, k int
	fmt.Fscanln(reader, &n, &k)

	var queue, result []int
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	var delIdx int
	for len(queue) > 0 {
		delIdx = (delIdx + k - 1) % len(queue)
		result = append(result, queue[delIdx])
		queue = append(queue[:delIdx], queue[delIdx+1:]...)
	}

	for i, v := range result {
		if i == 0 {
			fmt.Fprintf(writer, "<%d", v)
		} else {
			fmt.Fprintf(writer, ", %d", v)
		}
	}
	fmt.Fprintln(writer, ">")
}
