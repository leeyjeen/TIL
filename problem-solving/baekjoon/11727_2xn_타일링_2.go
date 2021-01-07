// https://www.acmicpc.net/problem/11727
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

	var count = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i == 1 {
			count[i] = 1
		} else if i == 2 {
			count[i] = 3
		} else {
			count[i] = (count[i-1] + 2*count[i-2]) % 10007
		}
	}
	fmt.Fprintln(writer, count[n])
}
