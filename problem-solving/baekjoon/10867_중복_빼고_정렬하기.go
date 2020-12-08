// https://www.acmicpc.net/problem/10867
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

	var n int
	fmt.Fscanln(reader, &n)

	var numbers = map[int]bool{}
	for i := 0; i < n; i++ {
		var key int
		fmt.Fscanf(reader, "%d ", &key)
		numbers[key] = true
	}

	keys := []int{}
	for key, _ := range numbers {
		keys = append(keys, key)
	}

	sort.Ints(keys)
	for _, key := range keys {
		fmt.Fprintf(writer, "%d ", key)
	}
	fmt.Fprintln(writer, "")
}
