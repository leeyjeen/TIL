// https://www.acmicpc.net/problem/1057
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

	var n, jimin, hansu int
	fmt.Fscanln(reader, &n, &jimin, &hansu)
	var round int
	for round < n {
		round++
		if jimin%2 == 0 {
			jimin /= 2
		} else {
			jimin = jimin/2 + 1
		}
		if hansu%2 == 0 {
			hansu /= 2
		} else {
			hansu = hansu/2 + 1
		}
		if jimin == hansu {
			fmt.Fprintln(writer, round)
			return
		}
	}
	fmt.Fprintln(writer, -1)
}
