// https://www.acmicpc.net/problem/2204
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for {
		var n int
		fmt.Fscanln(reader, &n)
		if n == 0 {
			break
		}
		var words []string
		for i := 0; i < n; i++ {
			var word string
			fmt.Fscanln(reader, &word)
			words = append(words, word)
		}
		sort.Slice(words, func(i, j int) bool {
			return strings.ToLower(words[i]) < strings.ToLower(words[j])
		})
		fmt.Fprintln(writer, words[0])
	}
}
