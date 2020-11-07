// https://www.acmicpc.net/problem/18258
// 큐의 개념을 익히고 실습하는 문제. 연산 당 시간 복잡도가 O(1)이어야 한다는 점에 유의.
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
	for i := 0; i < n; i++ {
		var command string
		var num int
		fmt.Fscanln(reader, &command, &num)

		switch command {
		case "push":
			queue = append(queue, num)
		case "pop":
			if len(queue) > 1 {
				fmt.Fprintln(writer, queue[0])
				queue = queue[1:]
			} else if len(queue) == 1 {
				fmt.Fprintln(writer, queue[0])
				queue = []int{}
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "size":
			fmt.Fprintln(writer, len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			if len(queue) > 0 {
				fmt.Fprintln(writer, queue[0])
			} else {
				fmt.Fprintln(writer, -1)
			}
		case "back":
			if len(queue) > 0 {
				fmt.Fprintln(writer, queue[len(queue)-1])
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}
