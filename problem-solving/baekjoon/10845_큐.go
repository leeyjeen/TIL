// https://www.acmicpc.net/problem/10845
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

	queue := []int{}
	for i := 0; i < n; i++ {
		var command string
		var number int
		fmt.Fscanln(reader, &command, &number)

		switch command {
		case "push":
			queue = append(queue, number)
		case "pop":
			output := -1
			if len(queue) > 0 {
				output = queue[0]
				queue = queue[1:]
			}
			fmt.Fprintln(writer, output)
		case "size":
			fmt.Fprintln(writer, len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Fprintln(writer, 1)
			} else {
				fmt.Fprintln(writer, 0)
			}
		case "front":
			output := -1
			if len(queue) > 0 {
				output = queue[0]
			}
			fmt.Fprintln(writer, output)
		case "back":
			output := -1
			if len(queue) > 0 {
				output = queue[len(queue)-1]
			}
			fmt.Fprintln(writer, output)
		}
	}
}
