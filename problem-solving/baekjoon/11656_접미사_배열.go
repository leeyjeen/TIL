// https://www.acmicpc.net/problem/11656
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

	var s string
	fmt.Fscanln(reader, &s)
	var suffixes []string
	for i := 0; i < len(s); i++ {
		suffixes = append(suffixes, s[i:])
	}
	sort.Strings(suffixes)
	for i := 0; i < len(suffixes); i++ {
		fmt.Fprintln(writer, suffixes[i])
	}
}
