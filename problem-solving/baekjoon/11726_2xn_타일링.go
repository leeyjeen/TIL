// https://www.acmicpc.net/problem/11726
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
		if i < 3 {
			count[i] = i
		} else {
			count[i] = (count[i-1] + count[i-2]) % 10007
		}
	}
	fmt.Fprintln(writer, count[n])
}
