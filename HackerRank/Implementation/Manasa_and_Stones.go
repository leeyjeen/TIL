// https://www.hackerrank.com/challenges/manasa-and-stones/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, n, a, b int
	fmt.Fscanln(reader, &t)
	for i := 0; i < t; i++ {
		fmt.Fscanln(reader, &n)
		fmt.Fscanln(reader, &a)
		fmt.Fscanln(reader, &b)
		for _, v := range stones(n, a, b) {
			fmt.Fprintf(writer, "%d ", v)
		}
		fmt.Fprintln(writer, "")
	}
}

func stones(n, a, b int) []int {
	var resultMap = map[int]bool{}
	var result []int
	for i := 0; i < n; i++ {
		j := n - 1 - i
		lastVal := a*i + b*j
		if _, ok := resultMap[lastVal]; !ok {
			resultMap[lastVal] = true
			result = append(result, lastVal)
		}
	}
	sort.Ints(result)
	return result
}
