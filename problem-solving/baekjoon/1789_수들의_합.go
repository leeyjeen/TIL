// https://www.acmicpc.net/problem/1789
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

	var s int
	fmt.Fscanln(reader, &s)

	var start, end int = 1, s
	var mid int
	var answer int
	for start <= end {
		mid = (start + end) / 2
		if mid*(mid+1)/2 <= s {
			answer = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	fmt.Fprintln(writer, answer)
}
