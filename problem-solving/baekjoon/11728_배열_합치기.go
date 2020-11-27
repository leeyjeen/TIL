// https://www.acmicpc.net/problem/11728
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

	array := []int{}
	for i := 0; i < n; i++ {
		var input int
		fmt.Fscanf(reader, "%d ", &input)
		array = append(array, input)
	}
	for i := 0; i < m; i++ {
		var input int
		fmt.Fscanf(reader, "%d ", &input)
		array = append(array, input)
	}

	sort.Ints(array)
	for i := 0; i < n+m-1; i++ {
		fmt.Fprintf(writer, "%d ", array[i])
	}
	fmt.Fprintf(writer, "%d\n", array[n+m-1])
}
