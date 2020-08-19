// https://www.hackerrank.com/challenges/repeated-string/problem
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

	var s string
	var n int
	fmt.Fscanln(reader, &s)
	fmt.Fscanln(reader, &n)
	fmt.Fprintln(writer, repeatedString(s, n))
}

func repeatedString(s string, n int) int {
	var count int
	for _, v := range s {
		if string(v) == "a" {
			count++
		}
	}
	count *= (n / len(s))
	for i := 0; i < n%len(s); i++ {
		if string(s[i]) == "a" {
			count++
		}
	}
	return count
}
