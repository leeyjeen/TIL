// https://www.acmicpc.net/problem/1966
// 큐의 개념이 응용된 문제
package main

import (
	"bufio"
	"fmt"
	"os"
)

type printValues struct {
	initIndex int
	value     int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(reader, &n, &m)

		var queue = make([]printValues, n)
		for j := 0; j < n; j++ {
			fmt.Fscanf(reader, "%d ", &queue[j].value)
			queue[j].initIndex = j
		}

		printCount := 0
		for len(queue) > 0 {
			printIndex := 0
			for i := 1; i < len(queue); i++ {
				if queue[0].value < queue[i].value {
					printIndex = i
					break
				}
			}
			if printIndex > 0 {
				queue = append(queue[1:], queue[0])
			} else {
				printCount++
				if queue[0].initIndex == m {
					fmt.Fprintln(writer, printCount)
				}
				queue = queue[1:]
			}
		}
	}
}
