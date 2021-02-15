// https://www.acmicpc.net/problem/20301
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

	var n, k, m int
	fmt.Fscanln(reader, &n, &k, &m)
	queue := []int{}
	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}
	selected := 0
	removed := []int{}
	turnRight := false
	for len(queue) > 0 {
		if len(removed)%m == 0 {
			turnRight = !turnRight
		}
		if turnRight {
			selected = (selected + k - 1) % len(queue)
		} else {
			selected = (selected - k) % len(queue)
			for selected < 0 {
				selected += len(queue)
			}
		}
		removed = append(removed, queue[selected])
		queue = append(queue[:selected], queue[selected+1:]...)
	}
	for i := 0; i < n; i++ {
		fmt.Fprintln(writer, removed[i])
	}
}
