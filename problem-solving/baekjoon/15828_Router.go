// https://www.acmicpc.net/problem/15828
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
	for {
		var input int
		fmt.Fscanln(reader, &input)
		if input == -1 {
			break
		} else if input == 0 {
			queue = queue[1:]
		} else {
			if len(queue) < n {
				queue = append(queue, input)
			}
		}
	}
	if len(queue) == 0 {
		fmt.Fprintln(writer, "empty")
	} else {
		for i := 0; i < len(queue); i++ {
			fmt.Fprintf(writer, "%d ", queue[i])
		}
	}
}
