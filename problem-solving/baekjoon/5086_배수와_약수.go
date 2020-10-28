// https://www.acmicpc.net/problem/5086
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

	for {
		var a, b int
		fmt.Fscanln(reader, &a, &b)

		if a == 0 && b == 0 {
			break
		}

		var resultStr string

		if isFactor(a, b) {
			resultStr = "factor"
		} else if isMultiple(a, b) {
			resultStr = "multiple"
		} else {
			resultStr = "neither"
		}

		fmt.Fprintln(writer, resultStr)
	}
}

func isFactor(a, b int) bool {
	if b%a == 0 {
		return true
	}
	return false
}

func isMultiple(a, b int) bool {
	if a%b == 0 {
		return true
	}
	return false
}
