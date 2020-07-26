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

	var b, n, m int
	fmt.Fscanln(reader, &b, &n, &m)
	var keyboard = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &keyboard[i])
	}
	var drives = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &drives[i])
	}
	sort.Ints(keyboard)
	sort.Ints(drives)
	var result = -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			total := keyboard[i] + drives[j]
			if total > b {
				break
			}
			if total > result {
				result = total
			}
		}
	}
	fmt.Fprintln(writer, result)
}
