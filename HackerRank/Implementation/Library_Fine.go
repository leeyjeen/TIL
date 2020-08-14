// https://www.hackerrank.com/challenges/library-fine/problem
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

	var d1, d2, m1, m2, y1, y2 int
	fmt.Fscanln(reader, &d1, &m1, &y1)
	fmt.Fscanln(reader, &d2, &m2, &y2)
	fmt.Fprintln(writer, libraryFine(d1, m1, y1, d2, m2, y2))
}

func libraryFine(d1, m1, y1, d2, m2, y2 int) int {
	if y1 < y2 || y1 == y2 && m1 < m2 || y1 == y2 && m1 == m2 && d1 <= d2 {
		return 0
	} else if y1 == y2 && m1 == m2 {
		return 15 * (d1 - d2)
	} else if y1 == y2 && m1 > m2 {
		return 500 * (m1 - m2)
	}
	return 10000
}
