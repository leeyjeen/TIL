// https://www.acmicpc.net/problem/9655
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
	fmt.Fprintln(writer, getWinner(n))
}

func getWinner(n int) string {
	if n%2 == 0 {
		return "CY"
	}
	return "SK"
}
