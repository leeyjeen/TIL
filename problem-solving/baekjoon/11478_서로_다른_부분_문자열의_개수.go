// https://www.acmicpc.net/problem/11478
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
	fmt.Fscanln(reader, &s)
	strs := map[string]bool{}
	for i := 0; i < len(s); i++ {
		for j := 1; i+j <= len(s); j++ {
			strs[s[i:i+j]] = true
		}
	}
	fmt.Fprintln(writer, len(strs))
}
