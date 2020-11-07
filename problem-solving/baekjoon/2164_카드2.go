// https://www.acmicpc.net/problem/2164
// 큐를 사용하여 동작을 구현하는 문제
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

	var queue []int
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	var discard bool = true
	for len(queue) > 1 {
		if discard {
			queue = queue[1:]
		} else {
			queue = append(queue[1:], queue[0])
		}
		discard = !discard
	}

	fmt.Fprintln(writer, queue[0])
}
