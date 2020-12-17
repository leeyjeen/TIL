// https://www.acmicpc.net/problem/1009
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

	var t int
	fmt.Fscanln(reader, &t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		result := 1
		for j := 0; j < b; j++ {
			result *= a
			result %= 10
		}
		if result == 0 {
			fmt.Fprintln(writer, 10)
		} else {
			fmt.Fprintln(writer, result)
		}
	}
}
