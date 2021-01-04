// https://www.acmicpc.net/problem/1024
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

	var n, l int
	fmt.Fscanln(reader, &n, &l)

	start := n/l - l/2
	var sum int
	for {
		if start < 0 {
			start++
		}
		sum = 0
		for i := start; i < start+l; i++ {
			sum += i
		}
		if sum < n {
			start++
		} else if sum == n {
			if start < 0 {
				fmt.Fprintln(writer, -1)
				return
			}
			break
		} else if sum > n {
			l++
			if l > 100 {
				fmt.Fprintln(writer, -1)
				return
			}
			start = n/l - l/2
		}
	}
	for i := start; i < start+l; i++ {
		fmt.Fprintf(writer, "%d ", i)
	}
	fmt.Fprintln(writer, "")
}
