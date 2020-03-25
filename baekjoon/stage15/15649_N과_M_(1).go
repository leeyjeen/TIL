package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n, &m)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var visited = make([]bool, n+1)
	var result = make([]int, m+1)

	sequence(writer, visited, result, 0, n, m)
}

func sequence(writer *bufio.Writer, visited []bool, result []int, index, n, m int) {
	if index == m { // if result array is full
		for i := 0; i < m; i++ {
			fmt.Fprintf(writer, "%d ", result[i])
		}
		fmt.Fprint(writer, "\n")
		return
	}
	for i := 1; i < n+1; i++ {
		if visited[i] == true {
			continue
		}
		result[index] = i
		visited[i] = true
		sequence(writer, visited, result, index+1, n, m)
		visited[i] = false
	}
}
