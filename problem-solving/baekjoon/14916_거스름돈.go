// https://www.acmicpc.net/problem/14916
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

	numOfFive, numOfTwo := 0, 0
	var isAvailable bool

	for n >= 0 {
		if n%5 == 0 {
			numOfFive = n / 5
			isAvailable = true
			break
		} else {
			n -= 2
			numOfTwo++
		}
	}

	if isAvailable {
		fmt.Fprintln(writer, numOfFive+numOfTwo)
	} else {
		fmt.Fprintln(writer, -1)
	}
}
