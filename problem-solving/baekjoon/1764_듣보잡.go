// https://www.acmicpc.net/problem/1764
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

	var n, m int
	fmt.Fscanln(reader, &n, &m)

	var dbj = map[string]int{}
	for i := 0; i < n; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		dbj[input]++
	}

	for i := 0; i < m; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		dbj[input]++
	}

	keys := []string{}
	for key, val := range dbj {
		if val == 2 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	fmt.Fprintln(writer, len(keys))
	for _, key := range keys {
		fmt.Fprintln(writer, key)
	}
}
