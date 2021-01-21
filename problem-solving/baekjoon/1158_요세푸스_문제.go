// https://www.acmicpc.net/problem/1158
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

	result := Josephus(n, k)
	for i := 0; i < len(result); i++ {
		if i == 0 {
			fmt.Fprintf(writer, "<%d", result[i])
		} else {
			fmt.Fprintf(writer, ", %d", result[i])
		}
	}
	fmt.Fprintln(writer, ">")
}

func Josephus(n, k int) (result []int) {
	queue := []int{}
	for i := 1; i <= n; i++ {
		queue = append(queue, i)
	}

	index := 0
	for len(queue) > 0 {
		index = (index + k - 1) % len(queue)
		result = append(result, queue[index])
		queue = append(queue[:index], queue[index+1:]...)
	}
	return result
}
